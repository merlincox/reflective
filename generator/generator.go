package generator

import (
	"fmt"
	"math"
	"reflect"
)

type Indication string

const (
	defMinStrLen          = 0
	defMaxStrLen          = 256
	defMinSliceLen        = 0
	defMaxSliceLen        = 256
	defMinMapLen          = 0
	defMaxMapLen          = 256
	defNilPointerChance   = 0.5
	defBooleanFalseChance = 0.5

	MapKey    Indication = "MapKey"
	MapValue  Indication = "MapValue"
	Real      Indication = "Real"
	Imaginary Indication = "Imaginary"
)

var defRunes = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type generator struct {
	minInt *int
	maxInt *int
	intFn  func(x ...Namable) (int, bool)

	minInt8 *int8
	maxInt8 *int8
	int8Fn  func(x ...Namable) (int8, bool)

	minInt16 *int16
	maxInt16 *int16
	int16Fn  func(x ...Namable) (int16, bool)

	minInt32 *int32
	maxInt32 *int32
	int32Fn  func(x ...Namable) (int32, bool)

	minInt64 *int64
	maxInt64 *int64
	int64Fn  func(x ...Namable) (int64, bool)

	minUint *uint
	maxUint *uint
	uintFn  func(x ...Namable) (uint, bool)

	minUint8 *uint8
	maxUint8 *uint8
	uint8Fn  func(x ...Namable) (uint8, bool)

	minUint16 *uint16
	maxUint16 *uint16
	uint16Fn  func(x ...Namable) (uint16, bool)

	minUint32 *uint32
	maxUint32 *uint32
	uint32Fn  func(x ...Namable) (uint32, bool)

	minUint64 *uint64
	maxUint64 *uint64
	uint64Fn  func(x ...Namable) (uint64, bool)

	minFloat32 *float32
	maxFloat32 *float32
	float32Fn  func(x ...Namable) (float32, bool)

	minFloat64 *float64
	maxFloat64 *float64
	float64Fn  func(x ...Namable) (float64, bool)

	minStrLen *uint
	maxStrLen *uint
	runes     []rune
	stringFn  func(x ...Namable) (string, bool)

	booleanFalseChance *float64
	boolFn             func(x ...Namable) (bool, bool)

	pointerNilChance *float64
	nilPointerFn     func(x ...Namable) (bool, bool)

	minSliceLen *uint
	maxSliceLen *uint
	sliceLenFn  func(x ...Namable) (int, bool)

	minMapLen *uint
	maxMapLen *uint
	mapLenFn  func(x ...Namable) (int, bool)

	rand PseudoRandom
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

// WithOption adds an option to a generator. This may be useful if you wish to use generator methods in a custom function.
func (g *generator) WithOption(option Option) *generator {
	g = option(g)
	return g
}

// WithOptions adds an option to a generator. This may be useful if you wish to use generator methods in a custom function.
func (g *generator) WithOptions(options ...Option) *generator {
	for _, o := range options {
		g = o(g)
	}

	return g
}

func (g *generator) genFloat32(x ...Namable) float32 {
	if g.float32Fn != nil {
		if out, ok := g.float32Fn(x...); ok {
			return out
		}
	}
	var min float32 = -math.MaxFloat32
	if g.minFloat32 != nil {
		min = *g.minFloat32
	}
	var max float32 = math.MaxFloat32
	if g.maxFloat32 != nil {
		max = *g.maxFloat32
	}
	var divisor float32 = 2.0

	return ((g.Float32() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) genFloat64(x ...Namable) float64 {
	if g.float64Fn != nil {
		if out, ok := g.float64Fn(x...); ok {
			return out
		}
	}
	min := -math.MaxFloat64
	if g.minFloat64 != nil {
		min = *g.minFloat64
	}
	max := math.MaxFloat64
	if g.maxFloat64 != nil {
		max = *g.maxFloat64
	}
	divisor := 2.0

	return ((g.Float64() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) genBool(x ...Namable) bool {
	if g.boolFn != nil {
		if out, ok := g.boolFn(x...); ok {
			return out
		}
	}
	chance := defBooleanFalseChance
	if g.booleanFalseChance != nil {
		chance = *g.booleanFalseChance
	}
	return g.Float64() >= chance
}

func (g *generator) genUseNilPointer(x ...Namable) bool {
	if g.nilPointerFn != nil {
		if out, ok := g.nilPointerFn(x...); ok {
			return out
		}
	}
	chance := defNilPointerChance
	if g.pointerNilChance != nil {
		chance = *g.pointerNilChance
	}
	return g.Float64() < chance
}

func (g *generator) genString(x ...Namable) string {
	if g.stringFn != nil {
		if out, ok := g.stringFn(x...); ok {
			return out
		}
	}
	max := defMaxStrLen
	if g.maxStrLen != nil {
		max = int(*g.maxStrLen)
	}
	if max == 0 {
		return ""
	}
	min := defMinStrLen
	if g.minStrLen != nil {
		min = int(*g.minStrLen)
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

func (g *generator) genSliceLen(x ...Namable) int {
	if g.sliceLenFn != nil {
		if out, ok := g.sliceLenFn(x...); ok {
			return out
		}
	}
	max := defMaxSliceLen
	if g.maxSliceLen != nil {
		max = int(*g.maxSliceLen)
	}
	if max == 0 {
		return 0
	}
	min := defMinSliceLen
	if g.minSliceLen != nil {
		min = int(*g.minSliceLen)
	}
	return g.Intn(max-min) + min
}

func (g *generator) genMapLen(x ...Namable) int {
	if g.mapLenFn != nil {
		if out, ok := g.mapLenFn(x...); ok {
			return out
		}
	}
	max := defMaxMapLen
	if g.maxMapLen != nil {
		max = int(*g.maxMapLen)
	}
	if max == 0 {
		return 0
	}
	min := defMinMapLen
	if g.minMapLen != nil {
		min = int(*g.minMapLen)
	}
	return g.Intn(max-min) + min
}

func (g *generator) genInt(x ...Namable) int {
	if g.intFn != nil {
		if out, ok := g.intFn(x...); ok {
			return out
		}
	}
	var min int64 = math.MinInt
	if g.minInt != nil {
		min = int64(*g.minInt)
	}
	var max int64 = math.MaxInt
	if g.maxInt != nil {
		max = int64(*g.maxInt)
	}
	return int(g.InclusiveInt64n(min, max))
}

func (g *generator) genInt8(x ...Namable) int8 {
	if g.int8Fn != nil {
		if out, ok := g.int8Fn(x...); ok {
			return out
		}
	}
	var min int32 = math.MinInt8
	if g.minInt8 != nil {
		min = int32(*g.minInt8)
	}
	var max int32 = math.MaxInt8
	if g.maxInt8 != nil {
		max = int32(*g.maxInt8)
	}
	return int8(g.InclusiveInt32n(min, max))
}

func (g *generator) genInt16(x ...Namable) int16 {
	if g.int16Fn != nil {
		if out, ok := g.int16Fn(x...); ok {
			return out
		}
	}
	var min int32 = math.MinInt16
	if g.minInt16 != nil {
		min = int32(*g.minInt16)
	}
	var max int32 = math.MaxInt16
	if g.maxInt16 != nil {
		max = int32(*g.maxInt16)
	}
	return int16(g.InclusiveInt32n(min, max))
}

func (g *generator) genInt32(x ...Namable) int32 {
	if g.int32Fn != nil {
		if out, ok := g.int32Fn(x...); ok {
			return out
		}
	}
	var min int32 = math.MinInt32
	if g.minInt32 != nil {
		min = int32(*g.minInt32)
	}
	var max int32 = math.MaxInt32
	if g.maxInt32 != nil {
		max = int32(*g.maxInt32)
	}
	return int32(g.InclusiveInt32n(min, max))
}

func (g *generator) genInt64(x ...Namable) int64 {
	if g.int64Fn != nil {
		if out, ok := g.int64Fn(x...); ok {
			return out
		}
	}
	var min int64 = math.MinInt64
	if g.minInt64 != nil {
		min = *g.minInt64
	}
	var max int64 = math.MaxInt64
	if g.maxInt64 != nil {
		max = *g.maxInt64
	}
	return g.InclusiveInt64n(min, max)
}

func (g *generator) genUint(x ...Namable) uint {
	if g.uintFn != nil {
		if out, ok := g.uintFn(x...); ok {
			return out
		}
	}
	var max uint64 = math.MaxUint
	if g.maxUint != nil {
		max = uint64(*g.maxUint)
	}
	if max == 0 {
		return 0
	}
	var min uint64 = 0
	if g.minUint != nil {
		min = uint64(*g.minUint)
	}
	return uint(g.InclusiveUint64n(min, max))
}

func (g *generator) genUint8(x ...Namable) uint8 {
	if g.uint8Fn != nil {
		if out, ok := g.uint8Fn(x...); ok {
			return out
		}
	}
	var max uint32 = math.MaxUint8
	if g.maxUint8 != nil {
		max = uint32(*g.maxUint8)
	}
	if max == 0 {
		return 0
	}
	var min uint32 = 0
	if g.minUint8 != nil {
		min = uint32(*g.minUint8)
	}
	return uint8(g.InclusiveUint32n(min, max))
}

func (g *generator) genUint16(x ...Namable) uint16 {
	if g.uint16Fn != nil {
		if out, ok := g.uint16Fn(x...); ok {
			return out
		}
	}
	var max uint32 = math.MaxUint16
	if g.maxUint16 != nil {
		max = uint32(*g.maxUint16)
	}
	if max == 0 {
		return 0
	}
	var min uint32 = 0
	if g.minUint16 != nil {
		min = uint32(*g.minUint16)
	}
	return uint16(g.InclusiveUint32n(min, max))
}

func (g *generator) genUint32(x ...Namable) uint32 {
	if g.uint32Fn != nil {
		if out, ok := g.uint32Fn(x...); ok {
			return out
		}
	}
	var max uint32 = math.MaxUint32
	if g.maxUint32 != nil {
		max = uint32(*g.maxUint32)
	}
	if max == 0 {
		return 0
	}
	var min uint32 = 0
	if g.minUint32 != nil {
		min = uint32(*g.minUint32)
	}
	return uint32(g.InclusiveUint32n(min, max))
}

func (g *generator) genUint64(x ...Namable) uint64 {
	if g.uint64Fn != nil {
		if out, ok := g.uint64Fn(x...); ok {
			return out
		}
	}
	var min uint64 = 0
	if g.minUint64 != nil {
		min = *g.minUint64
	}
	var max uint64 = math.MaxUint64
	if g.maxUint64 != nil {
		min = *g.maxUint64
	}
	return g.InclusiveUint64n(min, max)
}

// FillRandomly fills a data structure pseudo-randomly. The argument must be a pointer to the structure.
func (g *generator) FillRandomly(a any) error {

	value := reflect.ValueOf(a)
	if value.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument to FillRandomly to must be a pointer")
	}

	g.fill(value.Elem())
	return nil
}

// FillRandomlyByValue fills a data structure pseudo-randomly. The argument must be the reflect.Value of the structure.
func (g *generator) FillRandomlyByValue(value reflect.Value) error {
	if !value.CanSet() {
		return fmt.Errorf("the argument to FillRandomlyByValue must be able to be set")
	}

	g.fill(value)
	return nil
}

func (g *generator) fill(value reflect.Value, visited ...Namable) {
	if !value.CanSet() {
		return
	}

	currentType := value.Type()
	pushed := pushNamed(visited, currentType)

	switch value.Kind() {
	case reflect.Pointer:
		if g.genUseNilPointer(pushed...) {
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
		g.fill(value.Elem(), pushed...)

	case reflect.Bool:
		randBool := g.genBool(pushed...)
		value.SetBool(randBool)

	case reflect.Int:
		randInt := int64(g.genInt(pushed...))
		value.SetInt(randInt)

	case reflect.Int8:
		randInt := int64(g.genInt8(pushed...))
		value.SetInt(randInt)

	case reflect.Int16:
		randInt := int64(g.genInt16(pushed...))
		value.SetInt(randInt)

	case reflect.Int32:
		randInt := int64(g.genInt32(pushed...))
		value.SetInt(randInt)

	case reflect.Int64:
		randInt := g.genInt64(pushed...)
		value.SetInt(randInt)

	case reflect.Uint:
		randUint := uint64(g.genUint(pushed...))
		value.SetUint(randUint)

	case reflect.Uint8:
		randUint := uint64(g.genUint8(pushed...))
		value.SetUint(randUint)

	case reflect.Uint16:
		randUint := uint64(g.genUint16(pushed...))
		value.SetUint(randUint)

	case reflect.Uint32:
		randUint := uint64(g.genUint32(pushed...))
		value.SetUint(randUint)

	case reflect.Uint64:
		randUint := g.genUint64(pushed...)
		value.SetUint(randUint)

	case reflect.Float32:
		randFloat := float64(g.genFloat32(pushed...))
		value.SetFloat(randFloat)

	case reflect.Float64:
		randFloat := g.genFloat64(pushed...)
		value.SetFloat(randFloat)

	case reflect.Complex64:
		realIndicator := newIndicator(Real)
		repushed := pushNamed(pushed, realIndicator)
		r := float32(g.genFloat32(repushed...))
		imagIndicator := newIndicator(Imaginary)
		repushed = pushNamed(pushed, imagIndicator)
		i := float32(g.genFloat32(repushed...))
		value.SetComplex(complex128(complex(r, i)))

	case reflect.Complex128:
		realIndicator := newIndicator(Real)
		repushed := pushNamed(pushed, realIndicator)
		r := float64(g.genFloat32(repushed...))
		imagIndicator := newIndicator(Imaginary)
		repushed = pushNamed(pushed, imagIndicator)
		i := float64(g.genFloat32(repushed...))
		value.SetComplex(complex(r, i))

	case reflect.String:
		randStringVal := g.genString(pushed...)
		value.SetString(randStringVal)

	case reflect.Slice:
		elementType := currentType.Elem()
		size := g.genSliceLen(pushed...)
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			g.fill(newElement, pushed...)
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		value.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			g.fill(value.Index(i), pushed...)
		}

	case reflect.Map:
		mapVal := reflect.MakeMap(currentType)
		size := g.genMapLen(pushed...)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(currentType.Elem()))
			valIndicator := newIndicator(MapValue)
			repushed := pushNamed(pushed, valIndicator)
			g.fill(newElement, repushed...)
			newKey := reflect.Indirect(reflect.New(currentType.Key()))
			keyIndicator := newIndicator(MapKey)
			repushed = pushNamed(pushed, keyIndicator)
			g.fill(newKey, repushed...)
			mapVal.SetMapIndex(newKey, newElement)
		}
		value.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {

			fieldType := structField{field: currentType.Field(i)}
			repushed := pushNamed(pushed, fieldType)

			g.fill(value.Field(i), repushed...)
		}
	}
}
