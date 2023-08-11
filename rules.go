package reflective

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

type generator interface {
	Int63() int64
	Int63n(n int64) int64
	Uint64() uint64
	Float64() float64
}

type defRandom struct{}

func (d defRandom) Int63() int64 {
	return rand.Int63()
}

func (d defRandom) Int63n(n int64) int64 {
	return rand.Int63n(n)
}

func (d defRandom) Uint64() uint64 {
	return rand.Uint64()
}

func (d defRandom) Float64() float64 {
	return rand.Float64()
}

func NewRules() *Rules {
	return &Rules{
		Default: defaultBaseRules.WithGenerator(defRandom{}),
		TypeMap: make(map[reflect.Type]*BaseRules),
	}
}

func (r *Rules) fillRandomly(val reflect.Value) error {
	if !val.CanSet() {
		return nil
	}
	baseRules, ok := r.TypeMap[reflect.TypeOf(val)]
	if !ok {
		baseRules = r.Default
	}
	switch val.Kind() {

	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		return r.fillRandomly(val.Elem())

	case reflect.Bool:
		val.SetBool(baseRules.MakeBool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min64, max64 := int64(math.MinInt64), int64(math.MaxInt64)
		val.SetInt(randInt64(min64, max64))

	case reflect.Uint16, reflect.Uint32, reflect.Uint64:
		randUint := uint64(rand.Uint32())
		val.SetUint(randUint)

	case reflect.Float32, reflect.Float64:
		randFloat := float64(rand.Float32())
		val.SetFloat(randFloat)

	case reflect.String:
		val.SetString(baseRules.MakeString())

	case reflect.Slice:
		elementType := val.Type().Elem()
		size := 1 + rand.Intn(16)
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			if err := r.fillRandomly(newElement); err != nil {
				return err
			}
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		val.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if err := r.fillRandomly(val.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Map:
		mapType := val.Type()
		mapVal := reflect.MakeMap(mapType)
		size := 1 + rand.Intn(16)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(mapType.Elem()))
			if err := r.fillRandomly(newElement); err != nil {
				return err
			}
			newKey := reflect.Indirect(reflect.New(mapType.Key()))
			if err := r.fillRandomly(newKey); err != nil {
				return err
			}
			mapVal.SetMapIndex(newKey, newElement)
		}
		val.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if err := r.fillRandomly(val.Field(i)); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported kind: %s", val.Kind().String())
	}
	return nil
}

type Rules struct {
	TypeMap map[reflect.Type]*BaseRules
	Default *BaseRules
}

func (r *Rules) WithTypeRule(a any, rules *BaseRules) *Rules {
	r.TypeMap[reflect.TypeOf(a)] = rules
	return r
}

type BaseRules struct {
	generator generator
	Bool      *BoolRules
	Slice     *SliceRules
	Map       *MapRules
	String    *StringRules
	Pointer   *PointerRules
	Int       *IntRules
	Int8      *Int8Rules
	Int16     *Int16Rules
	Int32     *Int32Rules
	Int64     *Int64Rules
	Uint      *UintRules
	Uint8     *Uint8Rules
	Uint16    *Uint16Rules
	Uint32    *Uint32Rules
	Uint64    *Uint64Rules
	Float32   *Float32Rules
	Float64   *Float64Rules
}

func (b *BaseRules) WithGenerator(g generator) *BaseRules {
	b.generator = g
	return b
}

func (b *BaseRules) MakeBool() bool {
	if b.Bool != nil && b.Bool.Fn != nil {
		return b.Bool.Fn()
	}
	return b.generator.Int63()%2 == 0
}

type BoolRules struct {
	Fn func() bool
}

type SliceRules struct {
	Min int
	Max int
	Fn  func() int
}

type MapRules struct {
	Min int
	Max int
	Fn  func() int
}

func (b *BaseRules) MakeString() string {
	if b.String != nil && b.String.Fn != nil {
		return b.String.Fn()
	}
	min, max := defStringRules.Min, defStringRules.Max
	letters := defStringRules.Letters
	if b.String != nil {
		min, max = b.String.Min, b.String.Max
		letters = b.String.Letters
	}
	size := int(b.generator.Int63n(int64(max)-int64(min)) + int64(min))
	runes := make([]rune, size)
	for i := range runes {
		runes[i] = letters[b.generator.Int63n(int64(len(letters)))]
	}
	return string(runes)
}

type StringRules struct {
	Letters []rune
	Min     int
	Max     int
	Fn      func() string
}

var defStringRules = StringRules{
	Letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	Min:     1,
	Max:     8,
}

type PointerRules struct {
	NilPossible float64
	Fn          bool
}

type IntRules struct {
	N1 int
	N2 int
	Fn func() int
}

type Int8Rules struct {
	N1 int8
	N2 int8
	Fn func() int8
}

type Int16Rules struct {
	N1 int16
	N2 int16
	Fn func() int16
}

type Int32Rules struct {
	N1 int32
	N2 int32
	Fn func() int32
}

type Int64Rules struct {
	N1 int64
	N2 int64
	Fn func() int64
}

type UintRules struct {
	N1 uint
	N2 uint
	Fn func() uint
}

type Uint8Rules struct {
	N1 uint8
	N2 uint8
	Fn func() uint8
}

type Uint16Rules struct {
	N1 uint16
	N2 uint16
	Fn func() uint16
}

type Uint32Rules struct {
	N1 uint32
	N2 uint32
	Fn func() uint32
}

type Uint64Rules struct {
	N1 uint64
	N2 uint64
	Fn func() uint64
}

type Float32Rules struct {
	N1 float32
	N2 float32
	Fn func() float32
}

type Float64Rules struct {
	N1 float64
	N2 float64
	Fn func() float64
}

var defaultBaseRules = BaseRules{
	Slice: &SliceRules{
		Min: 1,
		Max: 16,
	},
	Map: &MapRules{
		Min: 1,
		Max: 16,
	},
	String: &StringRules{
		Letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		Min:     1,
		Max:     16,
	},
	Pointer: &PointerRules{},
	Int: &IntRules{
		N1: math.MinInt,
		N2: math.MaxInt,
	},
	Int8: &Int8Rules{
		N1: math.MinInt8,
		N2: math.MaxInt8,
	},
	Int16: &Int16Rules{
		N1: math.MinInt16,
		N2: math.MaxInt16,
	},
	Int32: &Int32Rules{
		N1: math.MinInt32,
		N2: math.MaxInt32,
	},
	Int64: &Int64Rules{
		N1: math.MinInt64,
		N2: math.MaxInt64,
	},
	Uint: &UintRules{
		N1: 0,
		N2: math.MaxUint,
	},
	Uint8: &Uint8Rules{
		N1: 0,
		N2: math.MaxUint8,
	},
	Uint16: &Uint16Rules{
		N1: 0,
		N2: math.MaxUint16,
	},
	Uint32: &Uint32Rules{
		N1: 0,
		N2: math.MaxUint32,
	},
	Uint64: &Uint64Rules{
		N1: 0,
		N2: math.MaxUint64,
	},
	Float32: &Float32Rules{
		N1: float32(math.MinInt),
		N2: float32(math.MaxInt),
	},
	Float64: &Float64Rules{
		N1: float64(math.MinInt),
		N2: float64(math.MaxInt),
	},
}
