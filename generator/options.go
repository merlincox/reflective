package generator

import "fmt"

func UsePseudoRandom(rand PseudoRandom) Option {
	return func(g *generator) (*generator, error) {
		g.rand = rand
		return g, nil
	}
}

func PointerNilChance(chance float64) Option {
	return func(g *generator) (*generator, error) {
		if chance < 0 || chance > 1 {
			return nil, fmt.Errorf("PointerNilChance: chance must be in range 0 to 1")
		}
		g.pointerNilChance = &chance
		return g, nil
	}
}

func PointerNilFn(fn func(t *Matcher) (bool, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.pointerNilFn = append(g.pointerNilFn, fn)
		return g, nil
	}
}

func BoolFalseChance(chance float64) Option {
	return func(g *generator) (*generator, error) {
		if chance < 0 || chance > 1 {
			return nil, fmt.Errorf("BoolFalseChance: chance must be in range 0 to 1")
		}
		g.booleanFalseChance = &chance
		return g, nil
	}
}

func BoolFalseFn(fn func(t *Matcher) (bool, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.boolFalseFn = append(g.boolFalseFn, fn)
		return g, nil
	}
}

func IntRange(min, max int) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("IntRange: min may not exceed max")
		}
		g.intSet.min = &min
		g.intSet.max = &max
		return g, nil
	}
}

func IntFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.intSet.fn = append(g.intSet.fn, fn)
		return g, nil
	}
}

func Int8Range(min, max int8) Option {
	return func(g *generator) (*generator, error) {
		if min > max {
			return nil, fmt.Errorf("Int8Range: min may not exceed max")
		}
		g.int8Set.min = &min
		g.int8Set.max = &max
		return g, nil
	}
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
		g.int16Set.min = &min
		g.int16Set.max = &max
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
		g.int32Set.min = &min
		g.int32Set.max = &max
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
		g.int64Set.min = &min
		g.int64Set.max = &max
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
		g.uintSet.min = &min
		g.uintSet.max = &max
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
		g.uint8Set.min = &min
		g.uint8Set.max = &max
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
		g.uint16Set.min = &min
		g.uint16Set.max = &max
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
		g.uint32Set.min = &min
		g.uint32Set.max = &max
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
		g.uint64Set.min = &min
		g.uint64Set.max = &max
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
		g.float32Set.min = &min
		g.float32Set.max = &max
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
		g.float64Set.min = &min
		g.float64Set.max = &max
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
		g.stringLenSet.min = &min
		g.stringLenSet.max = &max
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
		g.sliceLenSet.min = &min
		g.sliceLenSet.max = &max
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
		g.mapLenSet.min = &min
		g.mapLenSet.max = &max
		return g, nil
	}
}

func MapLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) (*generator, error) {
		g.mapLenSet.fn = append(g.mapLenSet.fn, fn)
		return g, nil
	}
}
