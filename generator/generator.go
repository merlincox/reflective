package generator

import (
	"fmt"
	"math"
	"reflect"

	"pgregory.net/rand"
)

type Indication string

const (
	defMinStrLen          = 0
	defMaxStrLen          = 256
	defMinSliceLen        = 0
	defMaxSliceLen        = 256
	defMinMapLen          = 0
	defMaxMapLen          = 256
	defPointerNilChance   = 0.5
	defBooleanFalseChance = 0.5

	MapKey    Indication = "MapKey"
	MapValue  Indication = "MapValue"
	Real      Indication = "Real"
	Imaginary Indication = "Imaginary"
)

var defRunes = []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Rand interface {
	Intn(int) int
	Int31() int32
	Int31n(int32) int32
	Uint64() uint64
	Uint64n(uint64) uint64
	Float32() float32
	Float64() float64
}

type Named interface {
	Name() string
	PkgPath() string
	Kind() reflect.Kind
}

type structField struct {
	field reflect.StructField
}

func (w structField) Kind() reflect.Kind {
	return reflect.Invalid
}

func (w structField) PkgPath() string {
	return w.field.PkgPath
}

func (w structField) Name() string {
	return w.field.Name
}

type indicator struct {
	value Indication
}

func (w indicator) Kind() reflect.Kind {
	return reflect.Invalid
}

func (w indicator) PkgPath() string {
	return ""
}

func (w indicator) Name() string {
	return string(w.value)
}

func newIndicator(value Indication) indicator {
	return indicator{value: value}
}

func pushNamed(visited []Named, named Named) []Named {
	pushed := make([]Named, len(visited)+1)
	copy(pushed, visited)
	pushed[len(visited)] = named
	return pushed
}

type generator struct {
	minInt *int
	maxInt *int
	intFn  func(x ...Named) (int, bool)

	minInt8 *int8
	maxInt8 *int8
	int8Fn  func(x ...Named) (int8, bool)

	minInt16 *int16
	maxInt16 *int16
	int16Fn  func(x ...Named) (int16, bool)

	minInt32 *int32
	maxInt32 *int32
	int32Fn  func(x ...Named) (int32, bool)

	minInt64 *int64
	maxInt64 *int64
	int64Fn  func(x ...Named) (int64, bool)

	minUint *uint
	maxUint *uint
	uintFn  func(x ...Named) (uint, bool)

	minUint8 *uint8
	maxUint8 *uint8
	uint8Fn  func(x ...Named) (uint8, bool)

	minUint16 *uint16
	maxUint16 *uint16
	uint16Fn  func(x ...Named) (uint16, bool)

	minUint32 *uint32
	maxUint32 *uint32
	uint32Fn  func(x ...Named) (uint32, bool)

	minUint64 *uint64
	maxUint64 *uint64
	uint64Fn  func(x ...Named) (uint64, bool)

	minFloat32 *float32
	maxFloat32 *float32
	float32Fn  func(x ...Named) (float32, bool)

	minFloat64 *float64
	maxFloat64 *float64
	float64Fn  func(x ...Named) (float64, bool)

	minStrLen *uint
	maxStrLen *uint
	runes     []rune
	stringFn  func(x ...Named) (string, bool)

	booleanFalseChance *float64
	boolFn             func(x ...Named) (bool, bool)

	pointerNilChance *float64
	pointerNilFn     func(x ...Named) (bool, bool)

	minSliceLen *uint
	maxSliceLen *uint
	sliceLenFn  func(x ...Named) (int, bool)

	minMapLen *uint
	maxMapLen *uint
	mapLenFn  func(x ...Named) (int, bool)

	rand Rand
}

type Option func(*generator) *generator

func New(options ...Option) *generator {
	g := &generator{}

	for _, o := range options {
		g = o(g)
	}

	return g
}

func (g *generator) WithOption(option Option) *generator {
	g = option(g)
	return g
}

func UseRand(rand Rand) Option {
	return func(g *generator) *generator {
		g.rand = rand
		return g
	}
}

func PointerNilChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.pointerNilChance = &chance
		return g
	}
}

func PointerNilFn(fn func(x ...Named) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.pointerNilFn = fn
		return g
	}
}

func BooleanFalseChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.booleanFalseChance = &chance
		return g
	}
}

func BoolFn(fn func(x ...Named) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.boolFn = fn
		return g
	}
}

func IntRange(min, max int) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt = &min
		g.maxInt = &max
		return g
	}
}

func IntFn(fn func(x ...Named) (int, bool)) Option {
	return func(g *generator) *generator {
		g.intFn = fn
		return g
	}
}

func Int8Range(min, max int8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt8 = &min
		g.maxInt8 = &max
		return g
	}
}

func Int8Fn(fn func(x ...Named) (int8, bool)) Option {
	return func(g *generator) *generator {
		g.int8Fn = fn
		return g
	}
}

func Int16Range(min, max int16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt16 = &min
		g.maxInt16 = &max
		return g
	}
}

func Int16Fn(fn func(x ...Named) (int16, bool)) Option {
	return func(g *generator) *generator {
		g.int16Fn = fn
		return g
	}
}

func Int32Range(min, max int32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt32 = &min
		g.maxInt32 = &max
		return g
	}
}

func Int32Fn(fn func(x ...Named) (int32, bool)) Option {
	return func(g *generator) *generator {
		g.int32Fn = fn
		return g
	}
}

func Int64Range(min, max int64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt64 = &min
		g.maxInt64 = &max
		return g
	}
}

func Int64Fn(fn func(x ...Named) (int64, bool)) Option {
	return func(g *generator) *generator {
		g.int64Fn = fn
		return g
	}
}

func UintRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint = &min
		g.maxUint = &max
		return g
	}
}

func UintFn(fn func(x ...Named) (uint, bool)) Option {
	return func(g *generator) *generator {
		g.uintFn = fn
		return g
	}
}

func Uint8Range(min, max uint8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint8 = &min
		g.maxUint8 = &max
		return g
	}
}

func Uint8Fn(fn func(x ...Named) (uint8, bool)) Option {
	return func(g *generator) *generator {
		g.uint8Fn = fn
		return g
	}
}

func Uint16Range(min, max uint16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint16 = &min
		g.maxUint16 = &max
		return g
	}
}

func Uint16Fn(fn func(x ...Named) (uint16, bool)) Option {
	return func(g *generator) *generator {
		g.uint16Fn = fn
		return g
	}
}

func Uint32Range(min, max uint32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint32 = &min
		g.maxUint32 = &max
		return g
	}
}

func Uint32Fn(fn func(x ...Named) (uint32, bool)) Option {
	return func(g *generator) *generator {
		g.uint32Fn = fn
		return g
	}
}

func Uint64Range(min, max uint64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint64 = &min
		g.maxUint64 = &max
		return g
	}
}

func Uint64Fn(fn func(x ...Named) (uint64, bool)) Option {
	return func(g *generator) *generator {
		g.uint64Fn = fn
		return g
	}
}

func Float32Range(min, max float32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat32 = &min
		g.maxFloat32 = &max
		return g
	}
}

func Float32Fn(fn func(x ...Named) (float32, bool)) Option {
	return func(g *generator) *generator {
		g.float32Fn = fn
		return g
	}
}

func Float64Range(min, max float64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat64 = &min
		g.maxFloat64 = &max
		return g
	}
}

func Float64Fn(fn func(x ...Named) (float64, bool)) Option {
	return func(g *generator) *generator {
		g.float64Fn = fn
		return g
	}
}

func StringLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minStrLen = &min
		g.maxStrLen = &max
		return g
	}
}

func StringFn(fn func(x ...Named) (string, bool)) Option {
	return func(g *generator) *generator {
		g.stringFn = fn
		return g
	}
}

func Runes(runes []rune) Option {
	return func(g *generator) *generator {
		g.runes = runes
		return g
	}
}

func SliceLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minSliceLen = &min
		g.maxSliceLen = &max
		return g
	}
}

func SliceLenFn(fn func(x ...Named) (int, bool)) Option {
	return func(g *generator) *generator {
		g.sliceLenFn = fn
		return g
	}
}

func MapLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minMapLen = &min
		g.maxMapLen = &max
		return g
	}
}

func MapLenFn(fn func(x ...Named) (int, bool)) Option {
	return func(g *generator) *generator {
		g.mapLenFn = fn
		return g
	}
}

func (g *generator) Intn(i int) int {
	if g.rand != nil {
		return g.rand.Intn(i)
	}
	return rand.Intn(i)
}

func (g *generator) Int31() int32 {
	if g.rand != nil {
		return g.rand.Int31()
	}
	return rand.Int31()
}

func (g *generator) Int31n(i int32) int32 {
	if g.rand != nil {
		return g.rand.Int31n(i)
	}
	return rand.Int31n(i)
}

func (g *generator) Uint64() uint64 {
	if g.rand != nil {
		return g.rand.Uint64()
	}
	return rand.Uint64()
}

func (g *generator) Uint64n(u uint64) uint64 {
	if g.rand != nil {
		return g.rand.Uint64n(u)
	}
	return rand.Uint64n(u)
}

func (g *generator) Float32() float32 {
	if g.rand != nil {
		return g.rand.Float32()
	}
	return rand.Float32()
}

func (g *generator) Float64() float64 {
	if g.rand != nil {
		return g.rand.Float64()
	}
	return rand.Float64()
}

func (g *generator) useNilPointer(x ...Named) bool {
	if g.pointerNilFn != nil {
		if out, ok := g.pointerNilFn(x...); ok {
			return out
		}
	}
	chance := defPointerNilChance
	if g.pointerNilChance != nil {
		chance = *g.pointerNilChance
	}
	return g.Float64() < chance
}

func (g *generator) fillFloat32(x ...Named) float32 {
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

func (g *generator) fillFloat64(x ...Named) float64 {
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

func (g *generator) fillBool(x ...Named) bool {
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

func (g *generator) fillString(x ...Named) string {
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

func (g *generator) getSliceLen(x ...Named) int {
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

func (g *generator) getMapLen(x ...Named) int {
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

func (g *generator) fillInt(x ...Named) int {
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
	return int(g.rangeInt64(min, max))
}

func (g *generator) fillInt8(x ...Named) int8 {
	if g.int8Fn != nil {
		if out, ok := g.int8Fn(x...); ok {
			return out
		}
	}
	var min int64 = math.MinInt8
	if g.minInt8 != nil {
		min = int64(*g.minInt8)
	}
	var max int64 = math.MaxInt8
	if g.maxInt8 != nil {
		max = int64(*g.maxInt8)
	}
	return int8(g.rangeInt64(min, max))
}

func (g *generator) fillInt16(x ...Named) int16 {
	if g.int16Fn != nil {
		if out, ok := g.int16Fn(x...); ok {
			return out
		}
	}
	var min int64 = math.MinInt16
	if g.minInt16 != nil {
		min = int64(*g.minInt16)
	}
	var max int64 = math.MaxInt16
	if g.maxInt16 != nil {
		max = int64(*g.maxInt16)
	}
	return int16(g.rangeInt64(min, max))
}

func (g *generator) fillInt32(x ...Named) int32 {
	if g.int32Fn != nil {
		if out, ok := g.int32Fn(x...); ok {
			return out
		}
	}
	var min int64 = math.MinInt32
	if g.minInt32 != nil {
		min = int64(*g.minInt32)
	}
	var max int64 = math.MaxInt32
	if g.maxInt32 != nil {
		max = int64(*g.maxInt32)
	}
	return int32(g.rangeInt64(min, max))
}

func (g *generator) fillInt64(x ...Named) int64 {
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
	return g.rangeInt64(min, max)
}

func (g *generator) rangeInt64(min, max int64) int64 {
	umin, umax := mapI64ToU64(min), mapI64ToU64(max)
	return mapU64ToI64(g.rangeUint64(umin, umax))
}

func (g *generator) rangeUint64(min, max uint64) uint64 {
	if min == 0 && max == math.MaxUint64 {
		return g.Uint64()
	}
	return g.Uint64n(max-min) + min
}

func mapU64ToI64(n uint64) int64 {
	return int64(n - 1<<63)
}

func mapI64ToU64(n int64) uint64 {
	return uint64(n) + 1<<63
}

func (g *generator) fillUint(x ...Named) uint {
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
	return uint(g.rangeUint64(min, max))
}

func (g *generator) fillUint8(x ...Named) uint8 {
	if g.uint8Fn != nil {
		if out, ok := g.uint8Fn(x...); ok {
			return out
		}
	}
	var max uint64 = math.MaxUint8
	if g.maxUint8 != nil {
		max = uint64(*g.maxUint8)
	}
	if max == 0 {
		return 0
	}
	var min uint64 = 0
	if g.minUint8 != nil {
		min = uint64(*g.minUint8)
	}
	return uint8(g.rangeUint64(min, max))
}

func (g *generator) fillUint16(x ...Named) uint16 {
	if g.uint16Fn != nil {
		if out, ok := g.uint16Fn(x...); ok {
			return out
		}
	}
	var max uint64 = math.MaxUint16
	if g.maxUint16 != nil {
		max = uint64(*g.maxUint16)
	}
	if max == 0 {
		return 0
	}
	var min uint64 = 0
	if g.minUint16 != nil {
		min = uint64(*g.minUint16)
	}
	return uint16(g.rangeUint64(min, max))
}

func (g *generator) fillUint32(x ...Named) uint32 {
	if g.uint32Fn != nil {
		if out, ok := g.uint32Fn(x...); ok {
			return out
		}
	}
	var max uint64 = math.MaxUint32
	if g.maxUint32 != nil {
		max = uint64(*g.maxUint32)
	}
	if max == 0 {
		return 0
	}
	var min uint64 = 0
	if g.minUint32 != nil {
		min = uint64(*g.minUint32)
	}
	return uint32(g.rangeUint64(min, max))
}

func (g *generator) fillUint64(x ...Named) uint64 {
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
	return g.rangeUint64(min, max)
}

func (g *generator) FillRandomly(a any) error {
	val := reflect.ValueOf(a)
	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument to FillRandomly to must be a pointer")
	}

	g.fill(val.Elem())

	return nil
}

func (g *generator) fill(value reflect.Value, visited ...Named) {
	if !value.CanSet() {
		return
	}

	currentType := value.Type()
	pushed := pushNamed(visited, currentType)

	switch value.Kind() {
	case reflect.Ptr:
		if g.useNilPointer(pushed...) {
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
		randBool := g.fillBool(pushed...)
		value.SetBool(randBool)

	case reflect.Int:
		randInt := int64(g.fillInt(pushed...))
		value.SetInt(randInt)

	case reflect.Int8:
		randInt := int64(g.fillInt8(pushed...))
		value.SetInt(randInt)

	case reflect.Int16:
		randInt := int64(g.fillInt16(pushed...))
		value.SetInt(randInt)

	case reflect.Int32:
		randInt := int64(g.fillInt32(pushed...))
		value.SetInt(randInt)

	case reflect.Int64:
		randInt := g.fillInt64(pushed...)
		value.SetInt(randInt)

	case reflect.Uint:
		randUint := uint64(g.fillUint(pushed...))
		value.SetUint(randUint)

	case reflect.Uint8:
		randUint := uint64(g.fillUint8(pushed...))
		value.SetUint(randUint)

	case reflect.Uint16:
		randUint := uint64(g.fillUint16(pushed...))
		value.SetUint(randUint)

	case reflect.Uint32:
		randUint := uint64(g.fillUint32(pushed...))
		value.SetUint(randUint)

	case reflect.Uint64:
		randUint := g.fillUint64(pushed...)
		value.SetUint(randUint)

	case reflect.Float32:
		randFloat := float64(g.fillFloat32(pushed...))
		value.SetFloat(randFloat)

	case reflect.Float64:
		randFloat := g.fillFloat64(pushed...)
		value.SetFloat(randFloat)

	case reflect.Complex64:
		realIndicator := newIndicator(Real)
		repushed := pushNamed(pushed, realIndicator)
		r := float32(g.fillFloat32(repushed...))
		imagIndicator := newIndicator(Imaginary)
		repushed = pushNamed(pushed, imagIndicator)
		i := float32(g.fillFloat32(repushed...))
		value.SetComplex(complex128(complex(r, i)))

	case reflect.Complex128:
		realIndicator := newIndicator(Real)
		repushed := pushNamed(pushed, realIndicator)
		r := float64(g.fillFloat32(repushed...))
		imagIndicator := newIndicator(Imaginary)
		repushed = pushNamed(pushed, imagIndicator)
		i := float64(g.fillFloat32(repushed...))
		value.SetComplex(complex(r, i))

	case reflect.String:
		randStringVal := g.fillString(pushed...)
		value.SetString(randStringVal)

	case reflect.Slice:
		elementType := currentType.Elem()
		size := g.getSliceLen(pushed...)
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
		size := g.getMapLen(pushed...)
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
