package generator

import (
	"fmt"
	"math"
	"reflect"
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

var defRunes = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type generator struct {
	rand PseudoRandom

	runes              []rune
	stringFn           []func(t *Matcher) (string, bool)
	booleanFalseChance *float64
	boolFalseFn        []func(t *Matcher) (bool, bool)
	pointerNilChance   *float64
	pointerNilFn       []func(t *Matcher) (bool, bool)

	stringLenSet gen[int]
	mapLenSet    gen[int]
	sliceLenSet  gen[int]
	float32Set   gen[float32]
	float64Set   gen[float64]
	intSet       gen[int]
	int8Set      gen[int8]
	int16Set     gen[int16]
	int32Set     gen[int32]
	int64Set     gen[int64]
	uintSet      gen[uint]
	uint8Set     gen[uint8]
	uint16Set    gen[uint16]
	uint32Set    gen[uint32]
	uint64Set    gen[uint64]
}

type gen[T numeric] struct {
	min *T
	max *T
	fn  []func(t *Matcher) (T, T, bool)
}

// Option defines an option for customising the generator
type Option func(*generator) *generator

// New creates a new generator with zero or more Options
func New(options ...Option) *generator {
	g := &generator{}

	for _, o := range options {
		g = o(g)
	}

	return g
}

// WithOptions adds an option to a generator. This may be useful if you wish to use generator methods in a custom function.
func (g *generator) WithOptions(options ...Option) *generator {
	for _, o := range options {
		g = o(g)
	}

	return g
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
	strLen := g.Intn(max-min) + min
	runes := make([]rune, strLen)
	for i := range runes {
		if len(g.runes) != 0 {
			runes[i] = g.runes[g.Intn(len(g.runes))]
		} else {
			runes[i] = defRunes[g.Intn(len(defRunes))]
		}
	}
	return string(runes)
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
		return fmt.Errorf("the argument to FillRandomlyByValue must be able to be set")
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
			reflect.Interface,
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
	kind, kindType := g.getAnyKind(false)
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
		size := g.genSliceLen(t.forSliceLen(t.current))
		subType := kindType.Elem()
		sliceVal := reflect.MakeSlice(reflect.SliceOf(subType), 0, size)
		for i := 0; i < size; i++ {
			subVal := reflect.Indirect(reflect.New(subType))
			g.fill(subVal, t.forSlice(t.current, i))
			sliceVal = reflect.Append(sliceVal, subVal)
		}
		anyVal.Set(sliceVal)
	case reflect.Map:
		mapVal := reflect.MakeMap(kindType)
		size := g.genMapLen(t.forMapLen(t.current))
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(kindType.Elem()))
			g.fill(newElement, t.forMapValue(t.current))
			newKey := reflect.Indirect(reflect.New(kindType.Key()))
			g.fill(newKey, t.forMapKey(t.current))
			mapVal.SetMapIndex(newKey, newElement)
		}
		anyVal.Set(mapVal)
	default:
		return
	}

	return anyVal, true
}

func (g *generator) getAnyKind(key bool) (reflect.Kind, reflect.Type) {
	max := 30
	if key {
		max = 22
	}
	i := g.Intn(max)

	switch i {
	case 0, 1, 2, 3:
		return reflect.Bool, reflect.TypeOf(true)
	case 4, 5, 6, 7:
		return reflect.String, reflect.TypeOf("string")
	case 8:
		return reflect.Int, reflect.TypeOf(0)
	case 9:
		return reflect.Int8, reflect.TypeOf(int8(0))
	case 10:
		return reflect.Int16, reflect.TypeOf(int16(0))
	case 11:
		return reflect.Int32, reflect.TypeOf(int32(0))
	case 12:
		return reflect.Int64, reflect.TypeOf(int64(0))
	case 13:
		return reflect.Uint, reflect.TypeOf(uint(0))
	case 14:
		return reflect.Uint8, reflect.TypeOf(uint8(0))
	case 15:
		return reflect.Uint16, reflect.TypeOf(uint16(0))
	case 16:
		return reflect.Uint32, reflect.TypeOf(uint32(0))
	case 17:
		return reflect.Uint64, reflect.TypeOf(uint64(0))
	case 18, 19:
		return reflect.Float32, reflect.TypeOf(float32(0))
	case 20, 21:
		return reflect.Float64, reflect.TypeOf(float64(0))
	case 22, 23, 24, 25:
		_, subType := g.getAnyKind(false)
		return reflect.Slice, reflect.SliceOf(subType)
	case 26, 27, 28, 29:
		_, keyType := g.getAnyKind(true)
		_, valueType := g.getAnyKind(false)
		return reflect.Map, reflect.MapOf(keyType, valueType)
	}

	return reflect.Invalid, reflect.TypeOf(nil)
}

//Int
//Int8
//Int16
//Int32
//Int64
//Uint
//Uint8
//Uint16
//Uint32
//Uint64
//Uintptr
//Float32
//Float64
//Complex64
//Complex128
//Array
//Chan
//Func
//Interface
//Map
//Pointer
//Slice
//String
//Struct
