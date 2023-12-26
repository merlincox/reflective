package generator_test

import (
	"github.com/merlincox/reflective/generator"
	"github.com/stretchr/testify/assert"
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
}

func TestRanges(t *testing.T) {
	subject, _ := generator.New()
	subject, _ = subject.WithOptions(
		generator.IntRange(int(intval), int(intval)),
		generator.Int8Range(int8(int8val), int8(int8val)),
		//generator.Int16Range(int16(int16val), int16(int16val)),
		//generator.Int32Range(int32(int32val), int32(int32val)),
		//generator.Int64Range(int64(int64val), int64(int64val)),
		//generator.UintRange(uint(uintval), uint(uintval)),
		//generator.Uint8Range(uint8(uint8val), uint8(uint8val)),
		//generator.Uint16Range(uint16(uint16val), uint16(uint16val)),
		//generator.Uint32Range(uint32(uint32val), uint32(uint32val)),
		//generator.Uint64Range(uint64(uint64val), uint64(uint64val)),
		//generator.Float32Range(float32(float32val), float32(float32val)),
		//generator.Float64Range(float64(float64val), float64(float64val)),
		//generator.BoolTrueRatio(1),
		//generator.MapLenRange(int(mapLen), int(mapLen)),
		//generator.SliceLenRange(int(sliceLen), int(sliceLen)),
		//generator.PointerNilRatio(0),
		//generator.StringLenRange(int(stringLen), int(stringLen)),
	)

	e := new(Unmatched)
	subject.FillRandomly(e)

	assert.Equal(t, int(intval), e.Int)
	assert.Equal(t, int8(int8val), e.Int8)
	//assert.Equal(t, int16(int16val), e.Int16)
	//assert.Equal(t, int32(int32val), e.Int32)
	//assert.Equal(t, int64(int64val), e.Int64)
	//
	//assert.Equal(t, uint(uintval), e.Uint)
	//assert.Equal(t, uint8(uint8val), e.Uint8)
	//assert.Equal(t, uint16(uint16val), e.Uint16)
	//assert.Equal(t, uint32(uint32val), e.Uint32)
	//assert.Equal(t, uint64(uint64val), e.Uint64)
	//
	//assert.Equal(t, float32(float32val), e.Float32)
	//assert.Equal(t, float64(float64val), e.Float64)
	//
	//assert.Equal(t, true, e.Bool)
	//assert.Equal(t, int(mapLen), len(e.Map))
	//assert.Equal(t, int(sliceLen), len(e.Slice))
	////assert.Equal(t, int(stringLen), len(e.String))
	//assert.NotNil(t, e.Pointer)
}

func TestFns(t *testing.T) {
	subject, _ := generator.New()
	subject, _ = subject.WithOptions(
		generator.IntFn(func(m *generator.Matcher) (int, int, bool) {
			return int(intval), int(intval), true
		}),
		generator.Int8Fn(func(m *generator.Matcher) (int8, int8, bool) {
			return int8(int8val), int8(int8val), true
		}),
		generator.Int16Fn(func(m *generator.Matcher) (int16, int16, bool) {
			return int16(int16val), int16(int16val), true
		}),
		generator.Int32Fn(func(m *generator.Matcher) (int32, int32, bool) {
			return int32(int32val), int32(int32val), true
		}),
		generator.Int64Fn(func(m *generator.Matcher) (int64, int64, bool) {
			return int64(int64val), int64(int64val), true
		}),
		generator.UintFn(func(m *generator.Matcher) (uint, uint, bool) {
			return uint(uintval), uint(uintval), true
		}),
		generator.Uint8Fn(func(m *generator.Matcher) (uint8, uint8, bool) {
			return uint8(uint8val), uint8(uint8val), true
		}),
		generator.Uint16Fn(func(m *generator.Matcher) (uint16, uint16, bool) {
			return uint16(uint16val), uint16(uint16val), true
		}),
		generator.Uint32Fn(func(m *generator.Matcher) (uint32, uint32, bool) {
			return uint32(uint32val), uint32(uint32val), true
		}),
		generator.Uint64Fn(func(m *generator.Matcher) (uint64, uint64, bool) {
			return uint64(uint64val), uint64(uint64val), true
		}),
		generator.Float32Fn(func(m *generator.Matcher) (float32, float32, bool) {
			return float32(float32val), float32(float32val), true
		}),
		generator.Float64Fn(func(m *generator.Matcher) (float64, float64, bool) {
			return float64(float64val), float64(float64val), true
		}),
		generator.BoolTrueFn(func(t *generator.Matcher) (float64, bool) {
			return 1, true
		}),
		generator.MapLenFn(func(t *generator.Matcher) (int, int, bool) {
			return int(mapLen), int(mapLen), true
		}),
		generator.SliceLenFn(func(t *generator.Matcher) (int, int, bool) {
			return int(sliceLen), int(sliceLen), true
		}),
		generator.PointerNilFn(func(t *generator.Matcher) (float64, bool) {
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

	assert.Equal(t, true, u.Bool)
	assert.Equal(t, int(mapLen), len(u.Map))
	assert.Equal(t, int(sliceLen), len(u.Slice))
	assert.NotNil(t, u.Pointer)
}

func TestMatch(t *testing.T) {
	subject, _ := generator.New()
	subject, _ = subject.WithOptions(
		generator.IntFn(func(m *generator.Matcher) (int, int, bool) {
			if m.MatchesAFieldOf(Matched{}, "Int") {
				return int(intval2), int(intval2), true
			}
			if m.MatchesA(custint(0)) {
				return int(intval3), int(intval3), true
			}
			return int(intval), int(intval), true
		}),
	)

	h := new(Holder)
	subject.FillRandomly(h)
	u := h.Unmatched
	m := h.Matched

	assert.Equal(t, int(intval), u.Int)
	assert.Equal(t, int(intval2), m.Int)
	assert.Equal(t, int(intval), m.Int2)
	assert.Equal(t, int(intval3), int(m.CustInt))
}
