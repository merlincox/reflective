package generator

import (
	"fmt"
	"math"
	"reflect"

	"pgregory.net/rand"
)

const (
	defMinStrLen          = 0
	defMaxStrLen          = 256
	defMinSliceLen        = 0
	defMaxSliceLen        = 256
	defMinMapLen          = 0
	defMaxMapLen          = 256
	defPointerNilChance   = 0.5
	defBooleanFalseChance = 0.5
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

type generator struct {
	minInt *int
	maxInt *int
	intFn  func(x any) (int, bool)

	minInt8 *int8
	maxInt8 *int8
	int8Fn  func(x any) (int8, bool)

	minInt16 *int16
	maxInt16 *int16
	int16Fn  func(x any) (int16, bool)

	minInt32 *int32
	maxInt32 *int32
	int32Fn  func(x any) (int32, bool)

	minInt64 *int64
	maxInt64 *int64
	int64Fn  func(x any) (int64, bool)

	minUint *uint
	maxUint *uint
	uintFn  func(x any) (uint, bool)

	minUint8 *uint8
	maxUint8 *uint8
	uint8Fn  func(x any) (uint8, bool)

	minUint16 *uint16
	maxUint16 *uint16
	uint16Fn  func(x any) (uint16, bool)

	minUint32 *uint32
	maxUint32 *uint32
	uint32Fn  func(x any) (uint32, bool)

	minUint64 *uint64
	maxUint64 *uint64
	uint64Fn  func(x any) (uint64, bool)

	minFloat32 *float32
	maxFloat32 *float32
	float32Fn  func(x any) (float32, bool)

	minFloat64 *float64
	maxFloat64 *float64
	float64Fn  func(x any) (float64, bool)

	minStrLen *uint
	maxStrLen *uint
	runes     []rune
	stringFn  func(x any) (string, bool)

	booleanFalseChance *float64
	BoolFn             func(x any) (bool, bool)

	pointerNilChance *float64
	PointerNilFn     func(x any) (bool, bool)

	minSliceLen *uint
	maxSliceLen *uint
	sliceLenFn  func(x any) (int, bool)

	minMapLen *uint
	maxMapLen *uint
	mapLenFn  func(x any) (int, bool)

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

func WithRand(rand Rand) Option {
	return func(g *generator) *generator {
		g.rand = rand
		return g
	}
}

//BoolFn   func(x any) (bool, bool)
//
//PointerNilFn func(x any) (bool, bool)

func WithPointerNilChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.pointerNilChance = &chance
		return g
	}
}

func WithBooleanFalseChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.booleanFalseChance = &chance
		return g
	}
}

func WithIntRange(min, max int) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt = &min
		g.maxInt = &max
		return g
	}
}

func WithIntFn(fn func(x any) (int, bool)) Option {
	return func(g *generator) *generator {
		g.intFn = fn
		return g
	}
}

func WithInt8Range(min, max int8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt8 = &min
		g.maxInt8 = &max
		return g
	}
}

func WithInt8Fn(fn func(x any) (int8, bool)) Option {
	return func(g *generator) *generator {
		g.int8Fn = fn
		return g
	}
}

func WithInt16Range(min, max int16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt16 = &min
		g.maxInt16 = &max
		return g
	}
}

func WithInt16Fn(fn func(x any) (int16, bool)) Option {
	return func(g *generator) *generator {
		g.int16Fn = fn
		return g
	}
}

func WithInt32Range(min, max int32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt32 = &min
		g.maxInt32 = &max
		return g
	}
}

func WithInt32Fn(fn func(x any) (int32, bool)) Option {
	return func(g *generator) *generator {
		g.int32Fn = fn
		return g
	}
}

func WithInt64Range(min, max int64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt64 = &min
		g.maxInt64 = &max
		return g
	}
}

func WithInt64Fn(fn func(x any) (int64, bool)) Option {
	return func(g *generator) *generator {
		g.int64Fn = fn
		return g
	}
}

func WithUintRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint = &min
		g.maxUint = &max
		return g
	}
}

func WithUintFn(fn func(x any) (uint, bool)) Option {
	return func(g *generator) *generator {
		g.uintFn = fn
		return g
	}
}

func WithUint8Range(min, max uint8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint8 = &min
		g.maxUint8 = &max
		return g
	}
}

func WithUint8Fn(fn func(x any) (uint8, bool)) Option {
	return func(g *generator) *generator {
		g.uint8Fn = fn
		return g
	}
}

func WithUint16Range(min, max uint16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint16 = &min
		g.maxUint16 = &max
		return g
	}
}

func WithUint16Fn(fn func(x any) (uint16, bool)) Option {
	return func(g *generator) *generator {
		g.uint16Fn = fn
		return g
	}
}

func WithUint32Range(min, max uint32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint32 = &min
		g.maxUint32 = &max
		return g
	}
}

func WithUint32Fn(fn func(x any) (uint32, bool)) Option {
	return func(g *generator) *generator {
		g.uint32Fn = fn
		return g
	}
}

func WithUint64Range(min, max uint64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint64 = &min
		g.maxUint64 = &max
		return g
	}
}

func WithUint64Fn(fn func(x any) (uint64, bool)) Option {
	return func(g *generator) *generator {
		g.uint64Fn = fn
		return g
	}
}

func WithFloat32Range(min, max float32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat32 = &min
		g.maxFloat32 = &max
		return g
	}
}

func WithFloat32Fn(fn func(x any) (float32, bool)) Option {
	return func(g *generator) *generator {
		g.float32Fn = fn
		return g
	}
}

func WithFloat64Range(min, max float64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat64 = &min
		g.maxFloat64 = &max
		return g
	}
}

func WithFloat64Fn(fn func(x any) (float64, bool)) Option {
	return func(g *generator) *generator {
		g.float64Fn = fn
		return g
	}
}

func WithStringLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minStrLen = &min
		g.maxStrLen = &max
		return g
	}
}

func WithStringFn(fn func(x any) (string, bool)) Option {
	return func(g *generator) *generator {
		g.stringFn = fn
		return g
	}
}

func WithRunes(runes []rune) Option {
	return func(g *generator) *generator {
		g.runes = runes
		return g
	}
}

func WithSliceLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minSliceLen = &min
		g.maxSliceLen = &max
		return g
	}
}

func WithSliceLenFn(fn func(x any) (int, bool)) Option {
	return func(g *generator) *generator {
		g.sliceLenFn = fn
		return g
	}
}

func WithMapLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minMapLen = &min
		g.maxMapLen = &max
		return g
	}
}

func WithMapLenFn(fn func(x any) (int, bool)) Option {
	return func(g *generator) *generator {
		g.mapLenFn = fn
		return g
	}
}

func (g *generator) randIntn(i int) int {
	if g.rand != nil {
		return g.rand.Intn(i)
	}
	return rand.Intn(i)
}

func (g *generator) randInt31() int32 {
	if g.rand != nil {
		return g.rand.Int31()
	}
	return rand.Int31()
}

func (g *generator) randInt31n(i int32) int32 {
	if g.rand != nil {
		return g.rand.Int31n(i)
	}
	return rand.Int31n(i)
}

func (g *generator) randUint64() uint64 {
	if g.rand != nil {
		return g.rand.Uint64()
	}
	return rand.Uint64()
}

func (g *generator) randUint64n(u uint64) uint64 {
	if g.rand != nil {
		return g.rand.Uint64n(u)
	}
	return rand.Uint64n(u)
}

func (g *generator) randFloat32() float32 {
	if g.rand != nil {
		return g.rand.Float32()
	}
	return rand.Float32()
}

func (g *generator) randFloat64() float64 {
	if g.rand != nil {
		return g.rand.Float64()
	}
	return rand.Float64()
}

func (g *generator) PointerNil(x any) bool {
	if g.PointerNilFn != nil {
		if out, ok := g.PointerNilFn(x); ok {
			return out
		}
	}
	chance := defPointerNilChance
	if g.pointerNilChance != nil {
		chance = *g.pointerNilChance
	}
	return g.randFloat64() < chance
}

func (g *generator) Float32(x any) float32 {
	if g.float32Fn != nil {
		if out, ok := g.float32Fn(x); ok {
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

	return ((g.randFloat32() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) Float64(x any) float64 {
	if g.float64Fn != nil {
		if out, ok := g.float64Fn(x); ok {
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

	return ((g.randFloat64() * ((max / divisor) - (min / divisor))) + (min / divisor)) * divisor
}

func (g *generator) String(x any) string {
	if g.stringFn != nil {
		if out, ok := g.stringFn(x); ok {
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
	strLen := g.randIntn(max-min) + min
	runes := make([]rune, strLen)
	for i := range runes {
		if len(g.runes) != 0 {
			runes[i] = g.runes[g.randIntn(len(g.runes))]
		} else {
			runes[i] = defRunes[g.randIntn(len(defRunes))]
		}
	}
	return string(runes)
}

func (g *generator) SliceLen(x any) int {
	if g.sliceLenFn != nil {
		if out, ok := g.sliceLenFn(x); ok {
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
	return g.randIntn(max-min) + min
}

func (g *generator) MapLen(x any) int {
	if g.mapLenFn != nil {
		if out, ok := g.mapLenFn(x); ok {
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
	return g.randIntn(max-min) + min
}

func (g *generator) Int(x any) int {
	if g.intFn != nil {
		if out, ok := g.intFn(x); ok {
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
	return int(g.genInt64(min, max))
}

func (g *generator) Int8(x any) int8 {
	if g.int8Fn != nil {
		if out, ok := g.int8Fn(x); ok {
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
	return int8(g.genInt64(min, max))
}

func (g *generator) Int16(x any) int16 {
	if g.int16Fn != nil {
		if out, ok := g.int16Fn(x); ok {
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
	return int16(g.genInt64(min, max))
}

func (g *generator) Int32(x any) int32 {
	if g.int32Fn != nil {
		if out, ok := g.int32Fn(x); ok {
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
	return int32(g.genInt64(min, max))
}

func (g *generator) Int64(x any) int64 {
	if g.int64Fn != nil {
		if out, ok := g.int64Fn(x); ok {
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
	return g.genInt64(min, max)
}

func (g *generator) genInt64(min, max int64) int64 {
	umin, umax := mapIToU(min), mapIToU(max)
	return mapUToI(g.genUint64(umin, umax))
}

func (g *generator) genUint64(min, max uint64) uint64 {
	if min == 0 && max == math.MaxUint64 {
		return g.randUint64()
	}
	return g.randUint64n(max-min) + min
}

func (g *generator) Uint(x any) uint {
	if g.uintFn != nil {
		if out, ok := g.uintFn(x); ok {
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
	return uint(g.genUint64(min, max))
}

func (g *generator) Uint8(x any) uint8 {
	if g.uint8Fn != nil {
		if out, ok := g.uint8Fn(x); ok {
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
	return uint8(g.genUint64(min, max))
}

func (g *generator) Uint16(x any) uint16 {
	if g.uint16Fn != nil {
		if out, ok := g.uint16Fn(x); ok {
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
	return uint16(g.genUint64(min, max))
}

func (g *generator) Uint32(x any) uint32 {
	if g.uint32Fn != nil {
		if out, ok := g.uint32Fn(x); ok {
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
	return uint32(g.genUint64(min, max))
}

func (g *generator) Uint64(x any) uint64 {
	if g.uint64Fn != nil {
		if out, ok := g.uint64Fn(x); ok {
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
	return g.genUint64(min, max)
}

func (g *generator) FillRandomly(a any) error {
	val := reflect.ValueOf(a)
	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument to FillRandomly to must be a pointer")
	}

	return g.fillRandomly(val.Elem())
}

func (g *generator) fillRandomly(receiver reflect.Value) error {
	if !receiver.CanSet() {
		return nil
	}

	switch receiver.Kind() {
	case reflect.Ptr:
		if g.PointerNil(receiver) {
			return nil
		}
		receiver.Set(reflect.New(receiver.Type().Elem()))
		return g.fillRandomly(receiver.Elem())

	case reflect.Bool:
		randBool := g.Bool(receiver)
		receiver.SetBool(randBool)

	case reflect.Int:
		randInt := int64(g.Int(receiver))
		receiver.SetInt(randInt)

	case reflect.Int8:
		randInt := int64(g.Int8(receiver))
		receiver.SetInt(randInt)

	case reflect.Int16:
		randInt := int64(g.Int16(receiver))
		receiver.SetInt(randInt)

	case reflect.Int32:
		randInt := int64(g.Int32(receiver))
		receiver.SetInt(randInt)

	case reflect.Int64:
		randInt := int64(g.Int64(receiver))
		receiver.SetInt(randInt)

	case reflect.Uint16, reflect.Uint32, reflect.Uint64:
		randUint := uint64(rand.Uint32())
		receiver.SetUint(randUint)

	case reflect.Float32, reflect.Float64:
		randFloat := float64(rand.Float32())
		receiver.SetFloat(randFloat)

	case reflect.String:
		randStringVal := g.String(receiver)
		receiver.SetString(randStringVal)

	case reflect.Slice:
		elementType := receiver.Type().Elem()
		size := g.SliceLen(receiver)
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			if err := g.fillRandomly(newElement); err != nil {
				return err
			}
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		receiver.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < receiver.Len(); i++ {
			if err := g.fillRandomly(receiver.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Map:
		mapType := receiver.Type()
		mapVal := reflect.MakeMap(mapType)
		size := g.MapLen(receiver)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(mapType.Elem()))
			if err := g.fillRandomly(newElement); err != nil {
				return err
			}
			newKey := reflect.Indirect(reflect.New(mapType.Key()))
			if err := g.fillRandomly(newKey); err != nil {
				return err
			}
			mapVal.SetMapIndex(newKey, newElement)
		}
		receiver.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < receiver.NumField(); i++ {
			if err := g.fillRandomly(receiver.Field(i)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported kind: %s", receiver.Kind().String())
	}
	return nil
}

func (g *generator) Bool(x any) bool {
	if g.BoolFn != nil {
		if out, ok := g.BoolFn(x); ok {
			return out
		}
	}
	chance := defBooleanFalseChance
	if g.booleanFalseChance != nil {
		chance = *g.booleanFalseChance
	}
	return g.randFloat64() >= chance

}

type Generator interface {
	Bool(x any) bool
	Int(x any) int
	Int8(x any) int8
	Int16(x any) int16
	Int32(x any) int32
	Int64(x any) int64
	Uint(x any) uint
	Uint8(x any) uint8
	Uint16(x any) uint16
	Uint32(x any) uint32
	Uint64(x any) uint64
	Float32(x any) float32
	Float64(x any) float64
	String(x any) string
	SliceLen(x any) int
	MapLen(x any) int
	PointerNotNil(x any) float64
	FillRandomly(a any) error
}

func mapUToI(n uint64) int64 {
	return int64(n - 1<<63)
}

func mapIToU(n int64) uint64 {
	if n >= 0 {
		return uint64(n) + 1<<63
	}
	return uint64(n + math.MaxInt64 + 1)
}
