package generator_test

import (
	"github.com/merlincox/reflective/generator"
	"pgregory.net/rand"
)

var _ generator.Randomiser = (*mockRand)(nil)

type mockRand struct {
	used int
}

func (m *mockRand) Uint32() uint32 {
	m.used++
	return rand.Uint32()
}

func (m *mockRand) Uint32n(u uint32) uint32 {
	m.used++
	return rand.Uint32n(u)
}

func (m *mockRand) Uint64() uint64 {
	m.used++
	return rand.Uint64()
}

func (m *mockRand) Uint64n(u uint64) uint64 {
	m.used++
	return rand.Uint64n(u)
}

func (m *mockRand) Float32() float32 {
	m.used++
	return rand.Float32()
}

func (m *mockRand) Float64() float64 {
	m.used++
	return rand.Float64()
}
