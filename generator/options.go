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

func PointerNilFn(fn func(x ...Namable) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.nilPointerFn = fn
		return g
	}
}

func BooleanFalseChance(chance float64) Option {
	if chance < 0 || chance > 1 {
		panic("chance must be in range 0 to 1")
	}
	return func(g *generator) *generator {
		g.booleanFalseChance = &chance
		return g
	}
}

func BoolFn(fn func(x ...Namable) (bool, bool)) Option {
	return func(g *generator) *generator {
		g.boolFn = fn
		return g
	}
}

func IntRange(min, max int) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt = &min
		g.maxInt = &max
		return g
	}
}

func IntFn(fn func(x ...Namable) (int, bool)) Option {
	return func(g *generator) *generator {
		g.intFn = fn
		return g
	}
}

func Int8Range(min, max int8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt8 = &min
		g.maxInt8 = &max
		return g
	}
}

func Int8Fn(fn func(x ...Namable) (int8, bool)) Option {
	return func(g *generator) *generator {
		g.int8Fn = fn
		return g
	}
}

func Int16Range(min, max int16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt16 = &min
		g.maxInt16 = &max
		return g
	}
}

func Int16Fn(fn func(x ...Namable) (int16, bool)) Option {
	return func(g *generator) *generator {
		g.int16Fn = fn
		return g
	}
}

func Int32Range(min, max int32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt32 = &min
		g.maxInt32 = &max
		return g
	}
}

func Int32Fn(fn func(x ...Namable) (int32, bool)) Option {
	return func(g *generator) *generator {
		g.int32Fn = fn
		return g
	}
}

func Int64Range(min, max int64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minInt64 = &min
		g.maxInt64 = &max
		return g
	}
}

func Int64Fn(fn func(x ...Namable) (int64, bool)) Option {
	return func(g *generator) *generator {
		g.int64Fn = fn
		return g
	}
}

func UintRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint = &min
		g.maxUint = &max
		return g
	}
}

func UintFn(fn func(x ...Namable) (uint, bool)) Option {
	return func(g *generator) *generator {
		g.uintFn = fn
		return g
	}
}

func Uint8Range(min, max uint8) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint8 = &min
		g.maxUint8 = &max
		return g
	}
}

func Uint8Fn(fn func(x ...Namable) (uint8, bool)) Option {
	return func(g *generator) *generator {
		g.uint8Fn = fn
		return g
	}
}

func Uint16Range(min, max uint16) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint16 = &min
		g.maxUint16 = &max
		return g
	}
}

func Uint16Fn(fn func(x ...Namable) (uint16, bool)) Option {
	return func(g *generator) *generator {
		g.uint16Fn = fn
		return g
	}
}

func Uint32Range(min, max uint32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint32 = &min
		g.maxUint32 = &max
		return g
	}
}

func Uint32Fn(fn func(x ...Namable) (uint32, bool)) Option {
	return func(g *generator) *generator {
		g.uint32Fn = fn
		return g
	}
}

func Uint64Range(min, max uint64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minUint64 = &min
		g.maxUint64 = &max
		return g
	}
}

func Uint64Fn(fn func(x ...Namable) (uint64, bool)) Option {
	return func(g *generator) *generator {
		g.uint64Fn = fn
		return g
	}
}

func Float32Range(min, max float32) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat32 = &min
		g.maxFloat32 = &max
		return g
	}
}

func Float32Fn(fn func(x ...Namable) (float32, bool)) Option {
	return func(g *generator) *generator {
		g.float32Fn = fn
		return g
	}
}

func Float64Range(min, max float64) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minFloat64 = &min
		g.maxFloat64 = &max
		return g
	}
}

func Float64Fn(fn func(x ...Namable) (float64, bool)) Option {
	return func(g *generator) *generator {
		g.float64Fn = fn
		return g
	}
}

func StringLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minStrLen = &min
		g.maxStrLen = &max
		return g
	}
}

func StringFn(fn func(x ...Namable) (string, bool)) Option {
	return func(g *generator) *generator {
		g.stringFn = fn
		return g
	}
}

func Runes(runes []rune) Option {
	return func(g *generator) *generator {
		g.runes = runes
		return g
	}
}

func SliceLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minSliceLen = &min
		g.maxSliceLen = &max
		return g
	}
}

func SliceLenFn(fn func(x ...Namable) (int, bool)) Option {
	return func(g *generator) *generator {
		g.sliceLenFn = fn
		return g
	}
}

func MapLenRange(min, max uint) Option {
	if min > max {
		panic("min may not exceed max")
	}
	return func(g *generator) *generator {
		g.minMapLen = &min
		g.maxMapLen = &max
		return g
	}
}

func MapLenFn(fn func(x ...Namable) (int, bool)) Option {
	return func(g *generator) *generator {
		g.mapLenFn = fn
		return g
	}
}
