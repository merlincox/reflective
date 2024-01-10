package generator

import (
	"fmt"
	"math"
)

// WithRandomiser replaces the default implementation of the Randomiser interface (pgregory.net/rand) with another.
func WithRandomiser(rand Randomiser) Option {
	return func(g *generator) (*generator, error) {
		g.rand = rand
		return g, nil
	}
}

// WithPointerNilRatio sets the probability of any pointer value being nil, where 0 means never and 1 means always
func WithPointerNilRatio(ratio float64) Option {
	return func(g *generator) (*generator, error) {
		if ratio < 0 || ratio > 1 {
			return nil, fmt.Errorf("WithPointerNilRatio: ratio must be in range 0 to 1")
		}
		g.pointerNilRatio = &ratio
		return g, nil
	}
}

// WithBoolTrueRatio sets the probability of any bool value being true, where 0 means never and 1 means always
func WithBoolTrueRatio(ratio float64) Option {
	return func(g *generator) (*generator, error) {
		if ratio < 0 || ratio > 1 {
			return nil, fmt.Errorf("WithBoolTrueRatio: ratio must be in range 0 to 1")
		}
		g.booleanTrueRatio = &ratio
		return g, nil
	}
}

func numericRange[T numeric](min, max T) Option {
	return func(g *generator) (*generator, error) {
		switch any(min).(type) {
		case stringLenInt, mapLenInt, sliceLenInt:
			if min < 0 {
				return nil, fmt.Errorf("length may not be negative")
			}
		case float32, float64:
			if math.IsNaN(float64(min)) || math.IsNaN(float64(max)) {
				return nil, fmt.Errorf("NaN is not supported")
			}
			if math.IsInf(float64(min), 1) || math.IsInf(float64(max), 1) {
				return nil, fmt.Errorf("infinity is not supported")
			}
			if math.IsInf(float64(min), -1) || math.IsInf(float64(max), -1) {
				return nil, fmt.Errorf("infinity is not supported")
			}
		}
		if min > max {
			return nil, fmt.Errorf("numeric range: min may not exceed max")
		}
		switch any(min).(type) {
		case int:
			g.intSet.interval = &interval[int]{min: int(min), max: int(max)}
		case stringLenInt:
			g.stringLenSet.interval = &interval[stringLenInt]{min: stringLenInt(min), max: stringLenInt(max)}
		case mapLenInt:
			g.mapLenSet.interval = &interval[mapLenInt]{min: mapLenInt(min), max: mapLenInt(max)}
		case sliceLenInt:
			g.sliceLenSet.interval = &interval[sliceLenInt]{min: sliceLenInt(min), max: sliceLenInt(max)}
		case int8:
			g.int8Set.interval = &interval[int8]{min: int8(min), max: int8(max)}
		case int16:
			g.int16Set.interval = &interval[int16]{min: int16(min), max: int16(max)}
		case int32:
			g.int32Set.interval = &interval[int32]{min: int32(min), max: int32(max)}
		case int64:
			g.int64Set.interval = &interval[int64]{min: int64(min), max: int64(max)}
		case uint:
			g.uintSet.interval = &interval[uint]{min: uint(min), max: uint(max)}
		case uint8:
			g.uint8Set.interval = &interval[uint8]{min: uint8(min), max: uint8(max)}
		case uint16:
			g.uint16Set.interval = &interval[uint16]{min: uint16(min), max: uint16(max)}
		case uint32:
			g.uint32Set.interval = &interval[uint32]{min: uint32(min), max: uint32(max)}
		case uint64:
			g.uint64Set.interval = &interval[uint64]{min: uint64(min), max: uint64(max)}
		case float32:
			g.float32Set.interval = &interval[float32]{min: float32(min), max: float32(max)}
		case float64:
			g.float64Set.interval = &interval[float64]{min: float64(min), max: float64(max)}
		}
		return g, nil
	}
}

func defaultInterval[T numeric]() interval[T] {
	var some T
	switch any(some).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return interval[T]{min: 0, max: T(defMaxInt)}
	case stringLenInt:
		return interval[T]{min: T(defMinStrLen), max: T(defMaxStrLen)}
	case mapLenInt:
		return interval[T]{min: T(defMinMapLen), max: T(defMaxMapLen)}
	case sliceLenInt:
		return interval[T]{min: T(defMinSliceLen), max: T(defMaxSliceLen)}
	}
	return interval[T]{min: 0, max: T(defMaxFloat)}
}

// WithRunes sets the runes from which strings are constructed
func WithRunes(runes []rune) Option {
	return func(g *generator) (*generator, error) {
		g.runes = runes
		return g, nil
	}
}

func WithStringLengthRange(min, max int) Option {
	return numericRange(stringLenInt(min), stringLenInt(max))
}

func WithSliceLengthRange(min, max int) Option {
	return numericRange(sliceLenInt(min), sliceLenInt(max))
}

func WithMapLengthRange(min, max int) Option {
	return numericRange(mapLenInt(min), mapLenInt(max))
}

func WithIntRange(min, max int) Option {
	return numericRange(min, max)
}

func WithInt8Range(min, max int8) Option {
	return numericRange(min, max)
}

func WithInt16Range(min, max int16) Option {
	return numericRange(min, max)
}

func WithInt32Range(min, max int32) Option {
	return numericRange(min, max)
}

func WithInt64Range(min, max int64) Option {
	return numericRange(min, max)
}

func WithUintRange(min, max uint) Option {
	return numericRange(min, max)
}

func WithUint8Range(min, max uint8) Option {
	return numericRange(min, max)
}

func WithUint16Range(min, max uint16) Option {
	return numericRange(min, max)
}

func WithUint32Range(min, max uint32) Option {
	return numericRange(min, max)
}

func WithUint64Range(min, max uint64) Option {
	return numericRange(min, max)
}

func WithFloat32Range(min, max float32) Option {
	return numericRange(min, max)
}

func WithFloat64Range(min, max float64) Option {
	return numericRange(min, max)
}

// WithPointerNilRatioFn registers a function for setting the chance of a pointer value being nil.
func WithPointerNilRatioFn(fn func(t *Matcher) (float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.pointerNilFns = append(g.pointerNilFns, fn)
		return g, nil
	}
}

// WithBoolTrueRatioFn registers a function for setting the chance of a boolean being true
func WithBoolTrueRatioFn(fn func(t *Matcher) (float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.boolTrueFns = append(g.boolTrueFns, fn)
		return g, nil
	}
}

// WithRunesFn registers a function for setting the runes from which strings are constructed within a matched context
func WithRunesFn(fn func(t *Matcher) ([]rune, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.runesFns = append(g.runesFns, fn)
		return g, nil
	}
}

func WithIntFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.intSet.fns = append(g.intSet.fns, fn)
		return g, nil
	}
}

func WithInt8Fn(fn func(t *Matcher) (int8, int8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int8Set.fns = append(g.int8Set.fns, fn)
		return g, nil
	}
}

func WithInt16Fn(fn func(t *Matcher) (int16, int16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int16Set.fns = append(g.int16Set.fns, fn)
		return g, nil
	}
}

func WithInt32Fn(fn func(t *Matcher) (int32, int32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int32Set.fns = append(g.int32Set.fns, fn)
		return g, nil
	}
}

func WithInt64Fn(fn func(t *Matcher) (int64, int64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int64Set.fns = append(g.int64Set.fns, fn)
		return g, nil
	}
}

func WithUintFn(fn func(t *Matcher) (uint, uint, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uintSet.fns = append(g.uintSet.fns, fn)
		return g, nil
	}
}

func WithUint8Fn(fn func(t *Matcher) (uint8, uint8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint8Set.fns = append(g.uint8Set.fns, fn)
		return g, nil
	}
}

func WithUint16Fn(fn func(t *Matcher) (uint16, uint16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint16Set.fns = append(g.uint16Set.fns, fn)
		return g, nil
	}
}

func WithUint32Fn(fn func(t *Matcher) (uint32, uint32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint32Set.fns = append(g.uint32Set.fns, fn)
		return g, nil
	}
}

func WithUint64Fn(fn func(t *Matcher) (uint64, uint64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint64Set.fns = append(g.uint64Set.fns, fn)
		return g, nil
	}
}

func WithFloat32Fn(fn func(t *Matcher) (float32, float32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float32Set.fns = append(g.float32Set.fns, fn)
		return g, nil
	}
}

func WithFloat64Fn(fn func(t *Matcher) (float64, float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float64Set.fns = append(g.float64Set.fns, fn)
		return g, nil
	}
}

func WithStringLengthFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (stringLenInt, stringLenInt, bool) {
		min, max, ok := fn(t)
		return stringLenInt(min), stringLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.stringLenSet.fns = append(g.stringLenSet.fns, adapter)
		return g, nil
	}
}

func WithSliceLengthFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (sliceLenInt, sliceLenInt, bool) {
		min, max, ok := fn(t)
		return sliceLenInt(min), sliceLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.sliceLenSet.fns = append(g.sliceLenSet.fns, adapter)
		return g, nil
	}
}

func WithMapLengthFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (mapLenInt, mapLenInt, bool) {
		min, max, ok := fn(t)
		return mapLenInt(min), mapLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.mapLenSet.fns = append(g.mapLenSet.fns, adapter)
		return g, nil
	}
}

func WithStringFn(fn func(t *Matcher) (string, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.stringFns = append(g.stringFns, fn)
		return g, nil
	}
}
