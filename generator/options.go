package generator

import (
	"fmt"
	"math"
)

func UsePseudoRandom(rand PseudoRandom) Option {
	return func(g *generator) (*generator, error) {
		g.rand = rand
		return g, nil
	}
}

func PointerNilRatio(ratio float64) Option {
	return func(g *generator) (*generator, error) {
		if ratio < 0 || ratio > 1 {
			return nil, fmt.Errorf("PointerNilRatio: ratio must be in range 0 to 1")
		}
		g.pointerNilRatio = &ratio
		return g, nil
	}
}

func PointerNilFn(fn func(t *Matcher) (float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.pointerNilFn = append(g.pointerNilFn, fn)
		return g, nil
	}
}

func BoolTrueRatio(ratio float64) Option {
	return func(g *generator) (*generator, error) {
		if ratio < 0 || ratio > 1 {
			return nil, fmt.Errorf("BoolTrueRatio: ratio must be in range 0 to 1")
		}
		g.booleanTrueRatio = &ratio
		return g, nil
	}
}

// BoolTrueFn registers a function for determining the chance of a boolean being true
func BoolTrueFn(fn func(t *Matcher) (float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.boolTrueFn = append(g.boolTrueFn, fn)
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
			g.intSet.minmax = &minmax[int]{min: int(min), max: int(max)}
		case stringLenInt:
			g.stringLenSet.minmax = &minmax[stringLenInt]{min: stringLenInt(min), max: stringLenInt(max)}
		case mapLenInt:
			g.mapLenSet.minmax = &minmax[mapLenInt]{min: mapLenInt(min), max: mapLenInt(max)}
		case sliceLenInt:
			g.sliceLenSet.minmax = &minmax[sliceLenInt]{min: sliceLenInt(min), max: sliceLenInt(max)}
		case int8:
			g.int8Set.minmax = &minmax[int8]{min: int8(min), max: int8(max)}
		case int16:
			g.int16Set.minmax = &minmax[int16]{min: int16(min), max: int16(max)}
		case int32:
			g.int32Set.minmax = &minmax[int32]{min: int32(min), max: int32(max)}
		case int64:
			g.int64Set.minmax = &minmax[int64]{min: int64(min), max: int64(max)}
		case uint:
			g.uintSet.minmax = &minmax[uint]{min: uint(min), max: uint(max)}
		case uint8:
			g.uint8Set.minmax = &minmax[uint8]{min: uint8(min), max: uint8(max)}
		case uint16:
			g.uint16Set.minmax = &minmax[uint16]{min: uint16(min), max: uint16(max)}
		case uint32:
			g.uint32Set.minmax = &minmax[uint32]{min: uint32(min), max: uint32(max)}
		case uint64:
			g.uint64Set.minmax = &minmax[uint64]{min: uint64(min), max: uint64(max)}
		case float32:
			g.float32Set.minmax = &minmax[float32]{min: float32(min), max: float32(max)}
		case float64:
			g.float64Set.minmax = &minmax[float64]{min: float64(min), max: float64(max)}
		}
		return g, nil
	}
}

func defaultMinMax[T numeric]() minmax[T] {
	var some T
	switch any(some).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return minmax[T]{min: 0, max: T(defMaxInt)}
	case stringLenInt:
		return minmax[T]{min: T(defMinStrLen), max: T(defMaxStrLen)}
	case mapLenInt:
		return minmax[T]{min: T(defMinMapLen), max: T(defMaxMapLen)}
	case sliceLenInt:
		return minmax[T]{min: T(defMinSliceLen), max: T(defMaxSliceLen)}
	}
	return minmax[T]{min: 0, max: T(defMaxFloat)}
}

func StringLenRange(min, max int) Option {
	return numericRange(stringLenInt(min), stringLenInt(max))
}

func SliceLenRange(min, max int) Option {
	return numericRange(sliceLenInt(min), sliceLenInt(max))
}

func MapLenRange(min, max int) Option {
	return numericRange(mapLenInt(min), mapLenInt(max))
}

func IntRange(min, max int) Option {
	return numericRange(min, max)
}

func Int8Range(min, max int8) Option {
	return numericRange(min, max)
}

func Int16Range(min, max int16) Option {
	return numericRange(min, max)
}

func Int32Range(min, max int32) Option {
	return numericRange(min, max)
}

func Int64Range(min, max int64) Option {
	return numericRange(min, max)
}

func UintRange(min, max uint) Option {
	return numericRange(min, max)
}

func Uint8Range(min, max uint8) Option {
	return numericRange(min, max)
}

func Uint16Range(min, max uint16) Option {
	return numericRange(min, max)
}

func Uint32Range(min, max uint32) Option {
	return numericRange(min, max)
}

func Uint64Range(min, max uint64) Option {
	return numericRange(min, max)
}

func Float32Range(min, max float32) Option {
	return numericRange(min, max)
}

func Float64Range(min, max float64) Option {
	return numericRange(min, max)
}

func IntFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.intSet.fn = append(g.intSet.fn, fn)
		return g, nil
	}
}

func Int8Fn(fn func(t *Matcher) (int8, int8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int8Set.fn = append(g.int8Set.fn, fn)
		return g, nil
	}
}

func Int16Fn(fn func(t *Matcher) (int16, int16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int16Set.fn = append(g.int16Set.fn, fn)
		return g, nil
	}
}

func Int32Fn(fn func(t *Matcher) (int32, int32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int32Set.fn = append(g.int32Set.fn, fn)
		return g, nil
	}
}

func Int64Fn(fn func(t *Matcher) (int64, int64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int64Set.fn = append(g.int64Set.fn, fn)
		return g, nil
	}
}

func UintFn(fn func(t *Matcher) (uint, uint, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uintSet.fn = append(g.uintSet.fn, fn)
		return g, nil
	}
}

func Uint8Fn(fn func(t *Matcher) (uint8, uint8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint8Set.fn = append(g.uint8Set.fn, fn)
		return g, nil
	}
}

func Uint16Fn(fn func(t *Matcher) (uint16, uint16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint16Set.fn = append(g.uint16Set.fn, fn)
		return g, nil
	}
}

func Uint32Fn(fn func(t *Matcher) (uint32, uint32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint32Set.fn = append(g.uint32Set.fn, fn)
		return g, nil
	}
}

func Uint64Fn(fn func(t *Matcher) (uint64, uint64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint64Set.fn = append(g.uint64Set.fn, fn)
		return g, nil
	}
}

func Float32Fn(fn func(t *Matcher) (float32, float32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float32Set.fn = append(g.float32Set.fn, fn)
		return g, nil
	}
}

func Float64Fn(fn func(t *Matcher) (float64, float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float64Set.fn = append(g.float64Set.fn, fn)
		return g, nil
	}
}

func StringLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (stringLenInt, stringLenInt, bool) {
		min, max, ok := fn(t)
		return stringLenInt(min), stringLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.stringLenSet.fn = append(g.stringLenSet.fn, adapter)
		return g, nil
	}
}

func SliceLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (sliceLenInt, sliceLenInt, bool) {
		min, max, ok := fn(t)
		return sliceLenInt(min), sliceLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.sliceLenSet.fn = append(g.sliceLenSet.fn, adapter)
		return g, nil
	}
}

func MapLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	adapter := func(t *Matcher) (mapLenInt, mapLenInt, bool) {
		min, max, ok := fn(t)
		return mapLenInt(min), mapLenInt(max), ok
	}

	return func(g *generator) (*generator, error) {
		g.mapLenSet.fn = append(g.mapLenSet.fn, adapter)
		return g, nil
	}
}

func StringFn(fn func(t *Matcher) (string, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.stringFn = append(g.stringFn, fn)
		return g, nil
	}
}

func Runes(runes []rune) Option {
	return func(g *generator) (*generator, error) {
		g.runes = runes
		return g, nil
	}
}
