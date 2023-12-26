package generator

import "fmt"

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

type stringLenInt int
type mapLenInt int
type sliceLenInt int

func numericRange[T numeric](min, max T) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("numeric range: min may not exceed max")
		}
		switch any(min).(type) {
		case int:
			g.intSet.minmax = &minmax[int]{min: int(min), max: int(max)}
		case stringLenInt:
			g.stringLenSet.minmax = &minmax[int]{min: int(min), max: int(max)}
		case mapLenInt:
			g.mapLenSet.minmax = &minmax[int]{min: int(min), max: int(max)}
		case sliceLenInt:
			g.sliceLenSet.minmax = &minmax[int]{min: int(min), max: int(max)}
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
	case float32, float64:
		return minmax[T]{min: 0, max: T(defMaxFloat)}
	}
	panic("unsupported numeric type")
}

func IntRange(min, max int) Option {
	return numericRange(min, max)
}

func IntFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.intSet.fn = append(g.intSet.fn, fn)
		return g, nil
	}
}

func Int8Range(min, max int8) Option {
	return numericRange(min, max)
}

func Int8Fn(fn func(t *Matcher) (int8, int8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int8Set.fn = append(g.int8Set.fn, fn)
		return g, nil
	}
}

func Int16Range(min, max int16) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Int16Range: min may not exceed max")
		}
		g.int16Set.minmax = &minmax[int16]{min: min, max: max}
		return g, nil
	}
}

func Int16Fn(fn func(t *Matcher) (int16, int16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int16Set.fn = append(g.int16Set.fn, fn)
		return g, nil
	}
}

func Int32Range(min, max int32) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Int32Range: min may not exceed max")
		}
		g.int32Set.minmax = &minmax[int32]{min: min, max: max}
		return g, nil
	}
}

func Int32Fn(fn func(t *Matcher) (int32, int32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int32Set.fn = append(g.int32Set.fn, fn)
		return g, nil
	}
}

func Int64Range(min, max int64) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Int64Range: min may not exceed max")
		}
		g.int64Set.minmax = &minmax[int64]{min: min, max: max}
		return g, nil
	}
}

func Int64Fn(fn func(t *Matcher) (int64, int64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.int64Set.fn = append(g.int64Set.fn, fn)
		return g, nil
	}
}

func UintRange(min, max uint) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("UintRange: min may not exceed max")
		}
		g.uintSet.minmax = &minmax[uint]{min: min, max: max}
		return g, nil
	}
}

func UintFn(fn func(t *Matcher) (uint, uint, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uintSet.fn = append(g.uintSet.fn, fn)
		return g, nil
	}
}

func Uint8Range(min, max uint8) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Uint8Range: min may not exceed max")
		}
		g.uint8Set.minmax = &minmax[uint8]{min: min, max: max}
		return g, nil
	}
}

func Uint8Fn(fn func(t *Matcher) (uint8, uint8, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint8Set.fn = append(g.uint8Set.fn, fn)
		return g, nil
	}
}

func Uint16Range(min, max uint16) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Uint16Range: min may not exceed max")
		}
		g.uint16Set.minmax = &minmax[uint16]{min: min, max: max}
		return g, nil
	}
}

func Uint16Fn(fn func(t *Matcher) (uint16, uint16, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint16Set.fn = append(g.uint16Set.fn, fn)
		return g, nil
	}
}

func Uint32Range(min, max uint32) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Uint32Range: min may not exceed max")
		}
		g.uint32Set.minmax = &minmax[uint32]{min: min, max: max}
		return g, nil
	}
}

func Uint32Fn(fn func(t *Matcher) (uint32, uint32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint32Set.fn = append(g.uint32Set.fn, fn)
		return g, nil
	}
}

func Uint64Range(min, max uint64) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Uint64Range: min may not exceed max")
		}
		g.uint64Set.minmax = &minmax[uint64]{min: min, max: max}
		return g, nil
	}
}

func Uint64Fn(fn func(t *Matcher) (uint64, uint64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.uint64Set.fn = append(g.uint64Set.fn, fn)
		return g, nil
	}
}

func Float32Range(min, max float32) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Float32Range: min may not exceed max")
		}
		g.float32Set.minmax = &minmax[float32]{min: min, max: max}
		return g, nil
	}
}

func Float32Fn(fn func(t *Matcher) (float32, float32, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float32Set.fn = append(g.float32Set.fn, fn)
		return g, nil
	}
}

func Float64Range(min, max float64) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Float64Range: min may not exceed max")
		}
		g.float64Set.minmax = &minmax[float64]{min: min, max: max}
		return g, nil
	}
}

func Float64Fn(fn func(t *Matcher) (float64, float64, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.float64Set.fn = append(g.float64Set.fn, fn)
		return g, nil
	}
}

func StringLenRange(min, max int) Option {
	return func(g *generator) (*generator, error) {
		if min < 0 {
			return nil, fmt.Errorf("StringLenRange: min may not be lower than zero")
		}
		if min > max {
			return nil, fmt.Errorf("StringLenRange: min may not exceed max")
		}
		g.stringLenSet.minmax = &minmax[int]{min: min, max: max}
		return g, nil
	}
}

func StringLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.stringLenSet.fn = append(g.stringLenSet.fn, fn)
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

func SliceLenRange(min, max int) Option {
	return func(g *generator) (*generator, error) {
		if min < 0 {
			return nil, fmt.Errorf("SliceLenRange: min may not be lower than zero")
		}
		if min > max {
			return nil, fmt.Errorf("SliceLenRange: min may not exceed max")
		}
		g.sliceLenSet.minmax = &minmax[int]{min: min, max: max}
		return g, nil
	}
}

func SliceLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.sliceLenSet.fn = append(g.sliceLenSet.fn, fn)
		return g, nil
	}
}

func MapLenRange(min, max int) Option {
	return func(g *generator) (*generator, error) {
		if min < 0 {
			return nil, fmt.Errorf("MapLenRange: min may not be lower than zero")
		}
		if min > max {
			return nil, fmt.Errorf("MapLenRange: min may not exceed max")
		}
		g.mapLenSet.minmax = &minmax[int]{min: min, max: max}
		return g, nil
	}
}

func MapLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.mapLenSet.fn = append(g.mapLenSet.fn, fn)
		return g, nil
	}
}
