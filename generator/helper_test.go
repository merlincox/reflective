package generator_test

import (
	"github.com/merlincox/reflective/generator"
	"pgregory.net/rand"
)

var _ generator.Randomiser = (*mockRand)(nil)

type mockRand struct {
	used bool
}

func (m *mockRand) Int() int {
	m.used = true
	return rand.Int()
}

func (m *mockRand) Intn(i int) int {
	m.used = true
	return rand.Intn(i)
}

func (m *mockRand) Int31() int32 {
	m.used = true
	return rand.Int31()
}

func (m *mockRand) Int31n(i int32) int32 {
	m.used = true
	return rand.Int31n(i)
}

func (m *mockRand) Int63() int64 {
	m.used = true
	return rand.Int63()
}

func (m *mockRand) Int63n(i int64) int64 {
	m.used = true
	return rand.Int63n(i)
}

func (m *mockRand) Uint32() uint32 {
	m.used = true
	return rand.Uint32()
}

func (m *mockRand) Uint32n(u uint32) uint32 {
	m.used = true
	return rand.Uint32n(u)
}

func (m *mockRand) Uint64() uint64 {
	m.used = true
	return rand.Uint64()
}

func (m *mockRand) Uint64n(u uint64) uint64 {
	m.used = true
	return rand.Uint64n(u)
}

func (m *mockRand) Float32() float32 {
	m.used = true
	return rand.Float32()
}

func (m *mockRand) Float64() float64 {
	m.used = true
	return rand.Float64()
}
