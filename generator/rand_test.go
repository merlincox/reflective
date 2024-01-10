package generator_test

import (
	"github.com/merlincox/reflective/generator"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGeneratorInclusiveInt32n(t *testing.T) {

	subject := generator.New()

	out := subject.InclusiveInt32n(5, 5)
	assert.Equal(t, int32(5), out)

	out = subject.InclusiveInt32n(math.MinInt32, math.MaxInt32)
	expected := out >= math.MinInt32 && out <= math.MaxInt32
	assert.True(t, expected)
}

func TestGeneratorInclusiveInt64n(t *testing.T) {

	subject := generator.New()

	out := subject.InclusiveInt64n(5, 5)
	assert.Equal(t, int64(5), out)

	out = subject.InclusiveInt64n(math.MinInt64, math.MaxInt64)
	expected := out >= math.MinInt64 && out <= math.MaxInt64
	assert.True(t, expected)
}

func TestGeneratorInclusiveUint32n(t *testing.T) {

	subject := generator.New()

	out := subject.InclusiveUint32n(5, 5)
	assert.Equal(t, uint32(5), out)

	out = subject.InclusiveUint32n(0, math.MaxUint32)
	expected := out >= 0 && out <= math.MaxUint32
	assert.True(t, expected)
}

func TestGeneratorInclusiveUint64n(t *testing.T) {

	subject := generator.New()

	out := subject.InclusiveUint64n(5, 5)
	assert.Equal(t, uint64(5), out)

	out = subject.InclusiveUint64n(0, math.MaxUint64)
	expected := out >= 0 && out <= math.MaxUint64
	assert.True(t, expected)
}
