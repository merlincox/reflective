package generator

func UsePseudoRandom(rand PseudoRandom) Option {
	return func(g *generator) *generator {
		g.rand = rand
		return g
	}
}

func PointerNilChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.pointerNilChance = &chance
		return g
	}
}

func PointerNilFn(fn func(t *Matcher) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.pointerNilFn = append(g.pointerNilFn, fn)
		return g
	}
}

func BoolFalseChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.booleanFalseChance = &chance
		return g
	}
}

func BoolFalseFn(fn func(t *Matcher) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.boolFalseFn = append(g.boolFalseFn, fn)
		return g
	}
}

func IntRange(min, max int) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.intSet.min = &min
		g.intSet.max = &max
		return g
	}
}

func IntFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) *generator {
		g.intSet.fn = append(g.intSet.fn, fn)
		return g
	}
}

func Int8Range(min, max int8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.int8Set.min = &min
		g.int8Set.max = &max
		return g
	}
}

func Int8Fn(fn func(t *Matcher) (int8, int8, bool)) Option {
	return func(g *generator) *generator {
		g.int8Set.fn = append(g.int8Set.fn, fn)
		return g
	}
}

func Int16Range(min, max int16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.int16Set.min = &min
		g.int16Set.max = &max
		return g
	}
}

func Int16Fn(fn func(t *Matcher) (int16, int16, bool)) Option {
	return func(g *generator) *generator {
		g.int16Set.fn = append(g.int16Set.fn, fn)
		return g
	}
}

func Int32Range(min, max int32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.int32Set.min = &min
		g.int32Set.max = &max
		return g
	}
}

func Int32Fn(fn func(t *Matcher) (int32, int32, bool)) Option {
	return func(g *generator) *generator {
		g.int32Set.fn = append(g.int32Set.fn, fn)
		return g
	}
}

func Int64Range(min, max int64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.int64Set.min = &min
		g.int64Set.max = &max
		return g
	}
}

func Int64Fn(fn func(t *Matcher) (int64, int64, bool)) Option {
	return func(g *generator) *generator {
		g.int64Set.fn = append(g.int64Set.fn, fn)
		return g
	}
}

func UintRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.uintSet.min = &min
		g.uintSet.max = &max
		return g
	}
}

func UintFn(fn func(t *Matcher) (uint, uint, bool)) Option {
	return func(g *generator) *generator {
		g.uintSet.fn = append(g.uintSet.fn, fn)
		return g
	}
}

func Uint8Range(min, max uint8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.uint8Set.min = &min
		g.uint8Set.max = &max
		return g
	}
}

func Uint8Fn(fn func(t *Matcher) (uint8, uint8, bool)) Option {
	return func(g *generator) *generator {
		g.uint8Set.fn = append(g.uint8Set.fn, fn)
		return g
	}
}

func Uint16Range(min, max uint16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.uint16Set.min = &min
		g.uint16Set.max = &max
		return g
	}
}

func Uint16Fn(fn func(t *Matcher) (uint16, uint16, bool)) Option {
	return func(g *generator) *generator {
		g.uint16Set.fn = append(g.uint16Set.fn, fn)
		return g
	}
}

func Uint32Range(min, max uint32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.uint32Set.min = &min
		g.uint32Set.max = &max
		return g
	}
}

func Uint32Fn(fn func(t *Matcher) (uint32, uint32, bool)) Option {
	return func(g *generator) *generator {
		g.uint32Set.fn = append(g.uint32Set.fn, fn)
		return g
	}
}

func Uint64Range(min, max uint64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.uint64Set.min = &min
		g.uint64Set.max = &max
		return g
	}
}

func Uint64Fn(fn func(t *Matcher) (uint64, uint64, bool)) Option {
	return func(g *generator) *generator {
		g.uint64Set.fn = append(g.uint64Set.fn, fn)
		return g
	}
}

func Float32Range(min, max float32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.float32Set.min = &min
		g.float32Set.max = &max
		return g
	}
}

func Float32Fn(fn func(t *Matcher) (float32, float32, bool)) Option {
	return func(g *generator) *generator {
		g.float32Set.fn = append(g.float32Set.fn, fn)
		return g
	}
}

func Float64Range(min, max float64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.float64Set.min = &min
		g.float64Set.max = &max
		return g
	}
}

func Float64Fn(fn func(t *Matcher) (float64, float64, bool)) Option {
	return func(g *generator) *generator {
		g.float64Set.fn = append(g.float64Set.fn, fn)
		return g
	}
}

func StringLenRange(min, max int) Option {
	if min < 0 {
		panic("min may not be lower than zero")
	}
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.stringLenSet.min = &min
		g.stringLenSet.max = &max
		return g
	}
}

func StringLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) *generator {
		g.stringLenSet.fn = append(g.stringLenSet.fn, fn)
		return g
	}
}

func StringFn(fn func(t *Matcher) (string, bool)) Option {
	return func(g *generator) *generator {
		g.stringFn = append(g.stringFn, fn)
		return g
	}
}

func Runes(runes []rune) Option {
	return func(g *generator) *generator {
		g.runes = runes
		return g
	}
}

func SliceLenRange(min, max int) Option {
	if min < 0 {
		panic("min may not be lower than zero")
	}
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.sliceLenSet.min = &min
		g.sliceLenSet.max = &max
		return g
	}
}

func SliceLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) *generator {
		g.sliceLenSet.fn = append(g.sliceLenSet.fn, fn)
		return g
	}
}

func MapLenRange(min, max int) Option {
	if min < 0 {
		panic("min may not be lower than zero")
	}
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.mapLenSet.min = &min
		g.mapLenSet.max = &max
		return g
	}
}

func MapLenFn(fn func(t *Matcher) (int, int, bool)) Option {
	return func(g *generator) *generator {
		g.mapLenSet.fn = append(g.mapLenSet.fn, fn)
		return g
	}
}
