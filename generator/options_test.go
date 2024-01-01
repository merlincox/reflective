package generator_test

import (
	"fmt"
	"github.com/merlincox/reflective/generator"
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

type inttest int64

const (
	noval inttest = iota
	intval
	int8val
	int16val
	int32val
	int64val
	uintval
	uint8val
	uint16val
	uint32val
	uint64val
	float32val
	float64val
	mapLen
	sliceLen
	stringLen

	intval2
	intval3
)

const runes = "abcdef"

type custint int

type Holder struct {
	Matched   Matched
	Unmatched Unmatched
}

type Unmatched struct {
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Float32 float32
	Float64 float64
	Bool    bool
	Map     map[string]string
	Slice   []string
	Pointer *string
	String  string
}

type Matched struct {
	Int     int
	Int2    int
	CustInt custint
	String1 string
	String2 string
}

func TestRanges(t *testing.T) {
	subject := generator.New()
	subject, _ = subject.WithOptions(
		generator.WithIntRange(int(intval), int(intval)),
		generator.WithInt8Range(int8(int8val), int8(int8val)),
		generator.WithInt16Range(int16(int16val), int16(int16val)),
		generator.WithInt32Range(int32(int32val), int32(int32val)),
		generator.WithInt64Range(int64(int64val), int64(int64val)),
		generator.WithUintRange(uint(uintval), uint(uintval)),
		generator.WithUint8Range(uint8(uint8val), uint8(uint8val)),
		generator.WithUint16Range(uint16(uint16val), uint16(uint16val)),
		generator.WithUint32Range(uint32(uint32val), uint32(uint32val)),
		generator.WithUint64Range(uint64(uint64val), uint64(uint64val)),
		generator.WithFloat32Range(float32(float32val), float32(float32val)),
		generator.WithFloat64Range(float64(float64val), float64(float64val)),
		generator.WithMapLengthRange(int(mapLen), int(mapLen)),
		generator.WithSliceLengthRange(int(sliceLen), int(sliceLen)),
		generator.WithStringLengthRange(int(stringLen), int(stringLen)),
		generator.WithBoolTrueRatio(1),
		generator.WithPointerNilRatio(0),
		generator.WithRunes([]rune(runes)),
	)

	e := new(Unmatched)
	subject.FillRandomly(e)

	assert.Equal(t, int(intval), e.Int)
	assert.Equal(t, int8(int8val), e.Int8)
	assert.Equal(t, int16(int16val), e.Int16)
	assert.Equal(t, int32(int32val), e.Int32)
	assert.Equal(t, int64(int64val), e.Int64)

	assert.Equal(t, uint(uintval), e.Uint)
	assert.Equal(t, uint8(uint8val), e.Uint8)
	assert.Equal(t, uint16(uint16val), e.Uint16)
	assert.Equal(t, uint32(uint32val), e.Uint32)
	assert.Equal(t, uint64(uint64val), e.Uint64)

	assert.Equal(t, float32(float32val), e.Float32)
	assert.Equal(t, float64(float64val), e.Float64)

	assert.Equal(t, int(mapLen), len(e.Map))
	assert.Equal(t, int(sliceLen), len(e.Slice))
	assert.Equal(t, int(stringLen), len(e.String))

	assert.Equal(t, true, e.Bool)
	assert.NotNil(t, e.Pointer)
	assert.True(t, strings.Contains(runes, e.String[0:1]))
}

func TestFns(t *testing.T) {
	subject := generator.New()
	subject, _ = subject.WithOptions(
		generator.WithIntFn(func(m *generator.Matcher) (int, int, bool) {
			return int(intval), int(intval), true
		}),
		generator.WithInt8Fn(func(m *generator.Matcher) (int8, int8, bool) {
			return int8(int8val), int8(int8val), true
		}),
		generator.WithInt16Fn(func(m *generator.Matcher) (int16, int16, bool) {
			return int16(int16val), int16(int16val), true
		}),
		generator.WithInt32Fn(func(m *generator.Matcher) (int32, int32, bool) {
			return int32(int32val), int32(int32val), true
		}),
		generator.WithInt64Fn(func(m *generator.Matcher) (int64, int64, bool) {
			return int64(int64val), int64(int64val), true
		}),
		generator.WithUintFn(func(m *generator.Matcher) (uint, uint, bool) {
			return uint(uintval), uint(uintval), true
		}),
		generator.WithUint8Fn(func(m *generator.Matcher) (uint8, uint8, bool) {
			return uint8(uint8val), uint8(uint8val), true
		}),
		generator.WithUint16Fn(func(m *generator.Matcher) (uint16, uint16, bool) {
			return uint16(uint16val), uint16(uint16val), true
		}),
		generator.WithUint32Fn(func(m *generator.Matcher) (uint32, uint32, bool) {
			return uint32(uint32val), uint32(uint32val), true
		}),
		generator.WithUint64Fn(func(m *generator.Matcher) (uint64, uint64, bool) {
			return uint64(uint64val), uint64(uint64val), true
		}),
		generator.WithFloat32Fn(func(m *generator.Matcher) (float32, float32, bool) {
			return float32(float32val), float32(float32val), true
		}),
		generator.WithFloat64Fn(func(m *generator.Matcher) (float64, float64, bool) {
			return float64(float64val), float64(float64val), true
		}),
		generator.WithMapLengthFn(func(t *generator.Matcher) (int, int, bool) {
			return int(mapLen), int(mapLen), true
		}),
		generator.WithSliceLengthFn(func(t *generator.Matcher) (int, int, bool) {
			return int(sliceLen), int(sliceLen), true
		}),
		generator.WithStringLengthFn(func(t *generator.Matcher) (int, int, bool) {
			return int(stringLen), int(stringLen), true
		}),
		generator.WithRunesFn(func(t *generator.Matcher) ([]rune, bool) {
			return []rune(runes), true
		}),
		generator.WithBoolTrueRatioFn(func(t *generator.Matcher) (float64, bool) {
			return 1, true
		}),
		generator.WithPointerNilRatioFn(func(t *generator.Matcher) (float64, bool) {
			return 0, true
		}),
	)

	u := new(Unmatched)
	subject.FillRandomly(u)

	assert.Equal(t, int(intval), u.Int)
	assert.Equal(t, int8(int8val), u.Int8)
	assert.Equal(t, int16(int16val), u.Int16)
	assert.Equal(t, int32(int32val), u.Int32)
	assert.Equal(t, int64(int64val), u.Int64)

	assert.Equal(t, uint(uintval), u.Uint)
	assert.Equal(t, uint8(uint8val), u.Uint8)
	assert.Equal(t, uint16(uint16val), u.Uint16)
	assert.Equal(t, uint32(uint32val), u.Uint32)
	assert.Equal(t, uint64(uint64val), u.Uint64)

	assert.Equal(t, float32(float32val), u.Float32)
	assert.Equal(t, float64(float64val), u.Float64)

	assert.Equal(t, int(mapLen), len(u.Map))
	assert.Equal(t, int(sliceLen), len(u.Slice))
	assert.Equal(t, int(stringLen), len(u.String))
	assert.True(t, strings.Contains(runes, u.String[0:1]))

	assert.Equal(t, true, u.Bool)
	assert.NotNil(t, u.Pointer)
}

func TestMatch(t *testing.T) {
	subject := generator.New()
	subject, _ = subject.WithOptions(
		generator.WithIntFn(func(m *generator.Matcher) (int, int, bool) {
			if m.MatchesAFieldOf(Matched{}, "Int") {
				return int(intval2), int(intval2), true
			}
			if m.MatchesA(custint(0)) {
				return int(intval3), int(intval3), true
			}
			return int(intval), int(intval), true
		}),
		generator.StringFn(func(m *generator.Matcher) (string, bool) {
			if m.MatchesAFieldOf(Matched{}, "String1") {
				return "TESTING", true
			}
			return "", false
		}),
	)

	h := new(Holder)
	_ = subject.FillRandomly(h)
	u := h.Unmatched
	m := h.Matched

	assert.Equal(t, int(intval), u.Int)
	assert.Equal(t, int(intval2), m.Int)
	assert.Equal(t, int(intval), m.Int2)
	assert.Equal(t, int(intval3), int(m.CustInt))
	assert.Equal(t, "TESTING", m.String1)
}

func TestRand(t *testing.T) {
	subject := generator.New()
	m := new(mockRand)
	subject, _ = subject.WithOptions(generator.WithRandomiser(m))
	_ = subject.Uint32()
	assert.True(t, m.used)
}

func TestErrors(t *testing.T) {
	type scenario struct {
		name        string
		option      generator.Option
		expectedErr error
	}
	scenarios := []scenario{
		{
			name:        "invalid bool true ratio",
			option:      generator.WithBoolTrueRatio(-1),
			expectedErr: fmt.Errorf("ratio must be in range 0 to 1"),
		},
		{
			name:        "invalid pointer nil ratio",
			option:      generator.WithPointerNilRatio(-1),
			expectedErr: fmt.Errorf("ratio must be in range 0 to 1"),
		},
		{
			name:        "invalid integer range",
			option:      generator.WithIntRange(5, 3),
			expectedErr: fmt.Errorf("numeric range: min may not exceed max"),
		},
		{
			name:        "invalid length",
			option:      generator.WithSliceLengthRange(-1, 3),
			expectedErr: fmt.Errorf("length may not be negative"),
		},
		{
			name:        "invalid float range with NaN",
			option:      generator.WithFloat64Range(math.NaN(), 1),
			expectedErr: fmt.Errorf("NaN is not supported"),
		},
		{
			name:        "invalid float range with infinity",
			option:      generator.WithFloat64Range(math.Inf(1), 1),
			expectedErr: fmt.Errorf("infinity is not supported"),
		},
		{
			name:        "invalid float range negative with infinity",
			option:      generator.WithFloat64Range(math.Inf(-1), 1),
			expectedErr: fmt.Errorf("infinity is not supported"),
		},
	}

	for _, s := range scenarios {
		s := s
		t.Run(s.name, func(tt *testing.T) {
			var err error
			subject := generator.New()
			subject, err = subject.WithOptions(s.option)
			assert.NotNil(t, err)
			assert.Contains(tt, err.Error(), s.expectedErr.Error())
		})
	}

}
