package generator

import (
	"math"

	"pgregory.net/rand"
)

var _ Randomiser = (*generator)(nil)

// Randomiser defines random methods.
type Randomiser interface {
	Int() int
	Intn(int) int
	Int31() int32
	Int31n(int32) int32
	Int63() int64
	Int63n(int64) int64
	Uint32() uint32
	Uint32n(uint32) uint32
	Uint64() uint64
	Uint64n(uint64) uint64
	Float32() float32
	Float64() float64
}

var _ inclusive = (*generator)(nil)

type inclusive interface {
	InclusiveInt32n(min, max int32) int32
	InclusiveUint32n(min, max uint32) uint32
	InclusiveInt64n(min, max int64) int64
	InclusiveUint64n(min, max uint64) uint64
}

// InclusiveInt32n returns a random int32 in the closed interval [min, max].
func (g *generator) InclusiveInt32n(min, max int32) int32 {
	if min == max {
		return max
	}
	return mapU32ToI32(g.InclusiveUint32n(mapI32ToU32(min), mapI32ToU32(max)))
}

// InclusiveUint32n returns a random uint32 in the closed interval [min, max].
func (g *generator) InclusiveUint32n(min, max uint32) uint32 {
	if min == max {
		return max
	}
	if min == 0 && max == math.MaxUint32 {
		return g.Uint32()
	}
	return g.Uint32n(max-min) + min
}

// InclusiveInt64n returns a random int64 in the closed interval [min, max].
func (g *generator) InclusiveInt64n(min, max int64) int64 {
	if min == max {
		return max
	}
	return mapU64ToI64(g.InclusiveUint64n(mapI64ToU64(min), mapI64ToU64(max)))
}

// InclusiveUint64n returns a random uint64 in the closed interval [min, max].
func (g *generator) InclusiveUint64n(min, max uint64) uint64 {
	if min == max {
		return max
	}
	if min == 0 && max == math.MaxUint64 {
		return g.Uint64()
	}
	return g.Uint64n(max-min) + min
}

// Int returns a uniformly distributed non-negative random int.
func (g *generator) Int() int {
	if g.rand != nil {
		return g.rand.Int()
	}
	return rand.Int()
}

// Intn returns, as an int, a uniformly distributed non-negative random number
// in the half-open interval [0, n). It panics if n <= 0.
func (g *generator) Intn(i int) int {
	if g.rand != nil {
		return g.rand.Intn(i)
	}
	return rand.Intn(i)
}

// Int31 returns a uniformly distributed non-negative random 31-bit integer as an int32.
func (g *generator) Int31() int32 {
	if g.rand != nil {
		return g.rand.Int31()
	}
	return rand.Int31()
}

// Int31n returns, as an int32, a uniformly distributed non-negative random number
// in the half-open interval [0, n). It panics if n <= 0.
func (g *generator) Int31n(i int32) int32 {
	if g.rand != nil {
		return g.rand.Int31n(i)
	}
	return rand.Int31n(i)
}

// Int63 returns a uniformly distributed non-negative random 63-bit integer as an int64.
func (g *generator) Int63() int64 {
	if g.rand != nil {
		return g.rand.Int63()
	}
	return rand.Int63()
}

// Int63n returns, as an int64, a uniformly distributed non-negative random number
// in the half-open interval [0, n). It panics if n <= 0.
func (g *generator) Int63n(i int64) int64 {
	if g.rand != nil {
		return g.rand.Int63n(i)
	}
	return rand.Int63n(i)
}

// Uint32 returns a uniformly distributed random 32-bit value as an uint32.
func (g *generator) Uint32() uint32 {
	if g.rand != nil {
		return g.rand.Uint32()
	}
	return rand.Uint32()
}

// Uint32n returns, as an uint32, a uniformly distributed random number in [0, n). Uint32n(0) returns 0.
func (g *generator) Uint32n(u uint32) uint32 {
	if g.rand != nil {
		return g.rand.Uint32n(u)
	}
	return rand.Uint32n(u)
}

// Uint64 returns a uniformly distributed random 64-bit value as an uint64.
func (g *generator) Uint64() uint64 {
	if g.rand != nil {
		return g.rand.Uint64()
	}
	return rand.Uint64()
}

// Uint64n returns, as an uint64, a uniformly distributed random number in [0, n). Uint64n(0) returns 0.
func (g *generator) Uint64n(u uint64) uint64 {
	if g.rand != nil {
		return g.rand.Uint64n(u)
	}
	return rand.Uint64n(u)
}

// Float32 returns, as a float32, a uniformly distributed random number in the half-open interval [0.0, 1.0).
func (g *generator) Float32() float32 {
	if g.rand != nil {
		return g.rand.Float32()
	}
	return rand.Float32()
}

// Float64 returns, as a float64, a uniformly distributed random number in the half-open interval [0.0, 1.0).
func (g *generator) Float64() float64 {
	if g.rand != nil {
		return g.rand.Float64()
	}
	return rand.Float64()
}

func mapU64ToI64(n uint64) int64 {
	return int64(n - 1<<63)
}

func mapI64ToU64(n int64) uint64 {
	return uint64(n) + 1<<63
}

func mapU32ToI32(n uint32) int32 {
	return int32(n - 1<<31)
}

func mapI32ToU32(n int32) uint32 {
	return uint32(n) + 1<<31
}
