package config

import (
	"math"
	"pgregory.net/rand"
	"reflect"
)

type Config struct {
	MinInt    *int
	MaxInt    *int
	IntFn     func(x any) int
	MinInt8   *int8
	MaxInt8   *int8
	Int8Fn    func() int8
	MinInt16  *int16
	MaxInt16  *int8
	Int16Fn   func() int16
	Int32Fn   func() int32
	Int64Fn   func() int64
	UintFn    func() uint
	Uint8Fn   func() uint8
	Uint16Fn  func() uint16
	Uint32Fn  func() uint32
	Uint64Fn  func() uint64
	Float32Fn func() float32
	Float64Fn func() float64

	StringFn  func() string
	PointerFn func() bool
	SliceFn   func() int
	MapFn     func() int
}

func (c Config) Int(x any) int {
	if c.IntFn != nil {
		return c.IntFn(x)
	}
	if c.MinInt == nil && c.MaxInt == nil {
		return rand.Int()
	}
	min := math.MinInt
	if c.MinInt != nil {
		min = *c.MinInt
	}
	max := math.MinInt
	if c.MaxInt != nil {
		min = *c.MaxInt
	}
	return rand.Intn(max-min) + max
}

func (c Config) Int8(x any) int8 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Int16(x any) int16 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Int32(x any) int32 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Int64(x any) int64 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Uint(x any) uint {
	//TODO implement me
	panic("implement me")
}

func (c Config) Uint8(x any) uint8 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Uint16(x any) uint16 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Uint32(x any) uint32 {
	//TODO implement me
	panic("implement me")
}

func (c Config) Uint64(x any) uint64 {
	//TODO implement me
	panic("implement me")
}

type Genxx interface {
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
}

type GeneratorMap map[reflect.Type]GeneratorMap

type G interface {
	Gen(a any) G
}

func ex() {

}
