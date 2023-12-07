package generator

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

const (
	defMinStrLen          = 4
	defMaxStrLen          = 16
	defMinSliceLen        = 2
	defMaxSliceLen        = 16
	defMinMapLen          = 2
	defMaxMapLen          = 16
	defNilPointerChance   = 0.5
	defBooleanFalseChance = 0.5
	defMaxInt             = int64(math.MaxInt8)
	defMaxFloat           = float64(math.MaxInt8)
)

var (
	defRunes = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	boolType    = reflect.TypeOf(true)
	intType     = reflect.TypeOf(0)
	int8Type    = reflect.TypeOf(int8(0))
	int16Type   = reflect.TypeOf(int16(0))
	int32Type   = reflect.TypeOf(int32(0))
	int64Type   = reflect.TypeOf(int64(0))
	uintType    = reflect.TypeOf(uint(0))
	uint8Type   = reflect.TypeOf(uint8(0))
	uint16Type  = reflect.TypeOf(uint16(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	uint64Type  = reflect.TypeOf(uint64(0))
	float32Type = reflect.TypeOf(float32(0))
	float64Type = reflect.TypeOf(float64(0))
	stringType  = reflect.TypeOf("")
)

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type generator struct {
	rand PseudoRandom

	maxDepth           int
	runes              []rune
	stringFn           []func(t *Matcher) (string, bool)
	booleanFalseChance *float64
	boolFalseFn        []func(t *Matcher) (bool, bool)
	pointerNilChance   *float64
	pointerNilFn       []func(t *Matcher) (bool, bool)

	stringLenSet nset[int]
	mapLenSet    nset[int]
	sliceLenSet  nset[int]
	float32Set   nset[float32]
	float64Set   nset[float64]
	intSet       nset[int]
	int8Set      nset[int8]
	int16Set     nset[int16]
	int32Set     nset[int32]
	int64Set     nset[int64]
	uintSet      nset[uint]
	uint8Set     nset[uint8]
	uint16Set    nset[uint16]
	uint32Set    nset[uint32]
	uint64Set    nset[uint64]
}

type nset[T numeric] struct {
	min *T
	max *T
	fn  []func(t *Matcher) (T, T, bool)
}

// Option defines an option for customising the generator
type Option func(*generator) (*generator, error)

// New creates a new generator with zero or more Options
func New(options ...Option) (*generator, error) {
	g := &generator{
		maxDepth: 3,
	}

	return g.WithOptions(options...)
}

// WithOptions adds options to a generator. This may be useful if you wish to use generator methods in a custom function.
func (g *generator) WithOptions(options ...Option) (*generator, error) {
	var err error
	for _, o := range options {
		g, err = o(g)
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

func (g *generator) genBool(t *Matcher) bool {
	for _, fn := range g.boolFalseFn {
		if out, ok := fn(t); ok {
			return out
		}
	}
	chance := defBooleanFalseChance
	if g.booleanFalseChance != nil {
		chance = *g.booleanFalseChance
	}
	return g.Float64() >= chance
}

func (g *generator) genUseNilPointer(t *Matcher) bool {
	for _, fn := range g.pointerNilFn {
		if out, ok := fn(t); ok {
			return out
		}
	}
	chance := defNilPointerChance
	if g.pointerNilChance != nil {
		chance = *g.pointerNilChance
	}
	return g.Float64() < chance
}

func (g *generator) genString(t *Matcher) string {
	for _, fn := range g.stringFn {
		if out, ok := fn(t); ok {
			return out
		}
	}
	set := g.stringLenSet
	var min, max = defMinStrLen, defMaxStrLen
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.max != nil {
			max = *set.max
		}
		if set.min != nil {
			max = *set.min
		}
	}
	if max == 0 {
		return ""
	}
	source := defRunes
	if len(g.runes) != 0 {
		source = g.runes
	}
	return g.fillString(g.Intn(max-min)+min, source)
}

func (g *generator) fillString(size int, source []rune) string {
	name := make([]rune, size)
	for j := 0; j < size; j++ {
		name[j] = source[g.Intn(len(source))]
	}
	return string(name)
}

func (g *generator) genFloat32(t *Matcher) float32 {
	set := g.float32Set
	var min, max float32 = 0.0, float32(defMaxFloat)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.min != nil {
			min = *set.min
		}
		if set.max != nil {
			max = *set.max
		}
	}
	var divisor float32 = 2.0

	return ((g.Float32() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) genFloat64(t *Matcher) float64 {
	set := g.float64Set
	var min, max = 0.0, defMaxFloat
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.min != nil {
			min = *set.min
		}
		if set.max != nil {
			max = *set.max
		}
	}
	var divisor = 2.0

	return ((g.Float64() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) genSliceLen(t *Matcher) int {
	set := g.sliceLenSet
	var min, max = defMinSliceLen, defMaxSliceLen
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.max != nil {
			max = *set.max
		}
		if set.min != nil {
			max = *set.min
		}
	}
	if max == 0 {
		return 0
	}
	return g.Intn(max-min) + min
}

func (g *generator) genMapLen(t *Matcher) int {
	set := g.mapLenSet
	var min, max = defMinMapLen, defMaxMapLen
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.max != nil {
			max = *set.max
		}
		if set.min != nil {
			max = *set.min
		}
	}
	if max == 0 {
		return 0
	}
	return g.Intn(max-min) + min
}

func (g *generator) genInt(t *Matcher) int {
	set := g.intSet
	var min, max int64 = 0, defMaxInt
	var found bool
	for _, fn := range g.intSet.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = int64(fmin), int64(fmax), true
		}
	}
	if !found {
		if set.min != nil {
			min = int64(*set.min)
		}
		if set.max != nil {
			max = int64(*set.max)
		}
	}
	return int(g.InclusiveInt64n(min, max))
}

func (g *generator) genInt8(t *Matcher) int8 {
	set := g.int8Set
	var min, max int32 = 0, int32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = int32(fmin), int32(fmax), true
		}
	}
	if !found {
		if set.min != nil {
			min = int32(*set.min)
		}
		if set.max != nil {
			max = int32(*set.max)
		}
	}
	return int8(g.InclusiveInt32n(min, max))
}

func (g *generator) genInt16(t *Matcher) int16 {
	set := g.int16Set
	var min, max int32 = 0, int32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = int32(fmin), int32(fmax), true
		}
	}
	if !found {
		if set.min != nil {
			min = int32(*set.min)
		}
		if set.max != nil {
			max = int32(*set.max)
		}
	}
	return int16(g.InclusiveInt32n(min, max))
}

func (g *generator) genInt32(t *Matcher) int32 {
	set := g.int32Set
	var min, max int32 = 0, int32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.min != nil {
			min = *set.min
		}
		if set.max != nil {
			max = *set.max
		}
	}
	return g.InclusiveInt32n(min, max)
}

func (g *generator) genInt64(t *Matcher) int64 {
	set := g.int64Set
	var min, max int64 = 0, defMaxInt
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if g.int64Set.min != nil {
			min = *set.min
		}
		if g.int64Set.max != nil {
			max = *set.max
		}
	}
	return g.InclusiveInt64n(min, max)
}

func (g *generator) genUint(t *Matcher) uint {
	set := g.uintSet
	var min, max uint64 = 0, uint64(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = uint64(fmin), uint64(fmax), true
		}
	}
	if !found {
		if g.uintSet.min != nil {
			min = uint64(*set.min)
		}
		if g.uintSet.max != nil {
			max = uint64(*set.max)
		}
	}
	return uint(g.InclusiveUint64n(min, max))
}

func (g *generator) genUint8(t *Matcher) uint8 {
	set := g.uint8Set
	var min, max uint32 = 0, uint32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = uint32(fmin), uint32(fmax), true
		}
	}
	if !found {
		if set.min != nil {
			min = uint32(*set.min)
		}
		if set.max != nil {
			max = uint32(*set.max)
		}
	}
	return uint8(g.InclusiveUint32n(min, max))
}

func (g *generator) genUint16(t *Matcher) uint16 {
	set := g.uint16Set
	var min, max uint32 = 0, uint32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = uint32(fmin), uint32(fmax), true
		}
	}
	if !found {
		if set.min != nil {
			min = uint32(*set.min)
		}
		if set.max != nil {
			max = uint32(*set.max)
		}
	}
	return uint16(g.InclusiveUint32n(min, max))
}

func (g *generator) genUint32(t *Matcher) uint32 {
	set := g.uint32Set
	var min, max uint32 = 0, uint32(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.min != nil {
			min = *set.min
		}
		if set.max != nil {
			max = *set.max
		}
	}
	return g.InclusiveUint32n(min, max)
}

func (g *generator) genUint64(t *Matcher) uint64 {
	set := g.uint64Set
	var min, max uint64 = 0, uint64(defMaxInt)
	var found bool
	for _, fn := range set.fn {
		if fmin, fmax, ok := fn(t); ok {
			if fmin == fmax {
				return fmin
			}
			min, max, found = fmin, fmax, true
		}
	}
	if !found {
		if set.min != nil {
			min = *set.min
		}
		if set.max != nil {
			max = *set.max
		}
	}

	return g.InclusiveUint64n(min, max)
}

// FillRandomly fills a data structure pseudo-randomly. The argument must be a pointer to the structure.
func (g *generator) FillRandomly(a any) error {

	value := reflect.ValueOf(a)
	if value.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument to FillRandomly to must be a pointer")
	}

	g.fill(value.Elem(), nil)
	return nil
}

// FillRandomlyByValue fills a data structure pseudo-randomly. The argument must be the reflect.Value of the structure.
func (g *generator) FillRandomlyByValue(value reflect.Value) error {
	if !value.CanSet() {
		return fmt.Errorf("the argument to FillRandomlyByValue must be able to be nset")
	}

	g.fill(value, nil)
	return nil
}

func (g *generator) fill(value reflect.Value, matcher *Matcher) {
	if !value.CanSet() {
		return
	}
	currentType := value.Type()

	switch value.Kind() {
	case reflect.Pointer:
		if g.genUseNilPointer(matcher.forType(currentType)) {
			return
		}
		pointerType := value.Type().Elem()
		switch pointerType.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Uintptr,
			reflect.UnsafePointer,
			reflect.Invalid:
			return
		}
		value.Set(reflect.New(pointerType))
		g.fill(value.Elem(), matcher.forType(currentType))

	case reflect.Bool:
		randBool := g.genBool(matcher.forType(currentType))
		value.SetBool(randBool)

	case reflect.Int:
		randInt := int64(g.genInt(matcher.forType(currentType)))
		value.SetInt(randInt)

	case reflect.Int8:
		randInt := int64(g.genInt8(matcher.forType(currentType)))
		value.SetInt(randInt)

	case reflect.Int16:
		randInt := int64(g.genInt16(matcher.forType(currentType)))
		value.SetInt(randInt)

	case reflect.Int32:
		randInt := int64(g.genInt32(matcher.forType(currentType)))
		value.SetInt(randInt)

	case reflect.Int64:
		randInt := g.genInt64(matcher.forType(currentType))
		value.SetInt(randInt)

	case reflect.Uint:
		randUint := uint64(g.genUint(matcher.forType(currentType)))
		value.SetUint(randUint)

	case reflect.Uint8:
		randUint := uint64(g.genUint8(matcher.forType(currentType)))
		value.SetUint(randUint)

	case reflect.Uint16:
		randUint := uint64(g.genUint16(matcher.forType(currentType)))
		value.SetUint(randUint)

	case reflect.Uint32:
		randUint := uint64(g.genUint32(matcher.forType(currentType)))
		value.SetUint(randUint)

	case reflect.Uint64:
		randUint := g.genUint64(matcher.forType(currentType))
		value.SetUint(randUint)

	case reflect.Float32:
		randFloat := float64(g.genFloat32(matcher.forType(currentType)))
		value.SetFloat(randFloat)

	case reflect.Float64:
		randFloat := g.genFloat64(matcher.forType(currentType))
		value.SetFloat(randFloat)

	case reflect.Complex64:
		r := float32(g.genFloat32(matcher.forReal(currentType)))
		i := float32(g.genFloat32(matcher.forImaginary(currentType)))
		value.SetComplex(complex128(complex(r, i)))

	case reflect.Complex128:
		r := float64(g.genFloat32(matcher.forReal(currentType)))
		i := float64(g.genFloat32(matcher.forImaginary(currentType)))
		value.SetComplex(complex(r, i))

	case reflect.String:
		randStringVal := g.genString(matcher.forType(currentType))
		value.SetString(randStringVal)

	case reflect.Interface:
		if currentType.NumMethod() == 0 {
			if newElement, ok := g.genAnyValue(matcher.forType(currentType)); ok {
				value.Set(newElement)
			}
		}

	case reflect.Slice:
		elementType := currentType.Elem()
		size := g.genSliceLen(matcher.forSliceLen(currentType))
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			g.fill(newElement, matcher.forSlice(currentType, i))
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		value.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			g.fill(value.Index(i), matcher.forSlice(currentType, i))
		}

	case reflect.Map:
		mapVal := reflect.MakeMap(currentType)
		size := g.genMapLen(matcher.forMapLen(currentType))
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(currentType.Elem()))
			g.fill(newElement, matcher.forMapValue(currentType))
			newKey := reflect.Indirect(reflect.New(currentType.Key()))
			g.fill(newKey, matcher.forMapKey(currentType))
			mapVal.SetMapIndex(newKey, newElement)
		}
		value.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			g.fill(value.Field(i), matcher.forField(currentType, currentType.Field(i)))
		}
	}
}

func (g *generator) genAnyValue(t *Matcher) (val reflect.Value, ok bool) {
	kind, kindType := g.pickAnyKind(false, 0)
	anyVal := reflect.Indirect(reflect.New(kindType))
	switch kind {
	case reflect.Bool:
		randVal := g.genBool(t)
		anyVal.SetBool(randVal)
	case reflect.String:
		randVal := g.genString(t)
		anyVal.SetString(randVal)
	case reflect.Int:
		randVal := g.genInt(t)
		anyVal.SetInt(int64(randVal))
	case reflect.Int8:
		randVal := g.genInt8(t)
		anyVal.SetInt(int64(randVal))
	case reflect.Int16:
		randVal := g.genInt16(t)
		anyVal.SetInt(int64(randVal))
	case reflect.Int32:
		randVal := g.genInt32(t)
		anyVal.SetInt(int64(randVal))
	case reflect.Int64:
		randVal := g.genInt64(t)
		anyVal.SetInt(randVal)
	case reflect.Uint:
		randVal := g.genUint(t)
		anyVal.SetUint(uint64(randVal))
	case reflect.Uint8:
		randVal := g.genUint8(t)
		anyVal.SetUint(uint64(randVal))
	case reflect.Uint16:
		randVal := g.genUint16(t)
		anyVal.SetUint(uint64(randVal))
	case reflect.Uint32:
		randVal := g.genUint32(t)
		anyVal.SetUint(uint64(randVal))
	case reflect.Uint64:
		randVal := g.genUint64(t)
		anyVal.SetUint(randVal)
	case reflect.Float32:
		randVal := g.genFloat32(t)
		anyVal.SetFloat(float64(randVal))
	case reflect.Float64:
		randVal := g.genFloat64(t)
		anyVal.SetFloat(randVal)
	case reflect.Slice:
		size := g.genSliceLen(t.forSliceLen(t.currentType))
		subType := kindType.Elem()
		sliceVal := reflect.MakeSlice(reflect.SliceOf(subType), 0, size)
		for i := 0; i < size; i++ {
			subVal := reflect.Indirect(reflect.New(subType))
			g.fill(subVal, t.forSlice(t.currentType, i))
			sliceVal = reflect.Append(sliceVal, subVal)
		}
		anyVal.Set(sliceVal)
	case reflect.Map:
		mapVal := reflect.MakeMap(kindType)
		size := g.genMapLen(t.forMapLen(t.currentType))
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(kindType.Elem()))
			g.fill(newElement, t.forMapValue(t.currentType))
			newKey := reflect.Indirect(reflect.New(kindType.Key()))
			g.fill(newKey, t.forMapKey(t.currentType))
			mapVal.SetMapIndex(newKey, newElement)
		}
		anyVal.Set(mapVal)
	case reflect.Struct:
		for i := 0; i < anyVal.NumField(); i++ {
			g.fill(anyVal.Field(i), nil)
		}

	default:
		return
	}

	return anyVal, true
}

func (g *generator) pickAnyKind(key bool, depth int) (reflect.Kind, reflect.Type) {
	max := 34
	if depth >= g.maxDepth {
		max = 22
	}
	if key {
		max = 14
	}
	i := g.Intn(max)

	switch i {
	case 0, 1, 2, 3:
		return reflect.String, stringType
	case 4:
		return reflect.Int, intType
	case 5:
		return reflect.Int8, int8Type
	case 6:
		return reflect.Int16, int16Type
	case 7:
		return reflect.Int32, int32Type
	case 8:
		return reflect.Int64, int64Type
	case 9:
		return reflect.Uint, uintType
	case 10:
		return reflect.Uint8, uint8Type
	case 11:
		return reflect.Uint16, uint16Type
	case 12:
		return reflect.Uint32, uint32Type
	case 13:
		return reflect.Uint64, uint64Type

		// non-keys
	case 14, 15:
		return reflect.Float32, float32Type
	case 16, 17:
		return reflect.Float64, float64Type
	case 18, 19, 20, 21:
		return reflect.Bool, boolType

		// depth-limited
	case 22, 23, 24, 25:
		//slice
		_, subType := g.pickAnyKind(false, depth+1)
		return reflect.Slice, reflect.SliceOf(subType)
	case 26, 27, 28, 29:
		//map
		_, keyType := g.pickAnyKind(true, depth+1)
		_, valueType := g.pickAnyKind(false, depth+1)
		return reflect.Map, reflect.MapOf(keyType, valueType)
	case 30, 31, 32, 33:
		//struct
		return reflect.Struct, reflect.StructOf(g.makeStructFields(depth))
	}

	return reflect.Invalid, reflect.TypeOf(nil)
}

func (g *generator) makeStructFields(depth int) []reflect.StructField {
	fields := make([]reflect.StructField, g.Intn(5)+1)
	runes := []rune("abcdefghijklmnopqrstuvwxyz")
	i := 0
	for {
		fieldName := strings.Title(g.fillString(g.Intn(5)+2, runes))
		dup := false
		for j := 0; j < i; j++ {
			if fieldName == fields[j].Name {
				dup = true
			}
		}
		if dup {
			continue
		}
		_, fieldType := g.pickAnyKind(false, depth+1)
		fields[i] = reflect.StructField{
			Name: fieldName,
			Type: fieldType,
		}
		i++
		if i == len(fields) {
			break
		}
	}
	return fields
}
