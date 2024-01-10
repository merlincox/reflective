package generator

import "math"

func (g *generator) chanceTrue(ratio float64) bool {
	if ratio <= 0 {
		return false
	}
	if ratio >= 1 {
		return true
	}
	return g.Float64() >= ratio
}

func (g *generator) genBool(t *Matcher) bool {
	for _, fn := range g.boolTrueFns {
		if out, ok := fn(t); ok {
			return g.chanceTrue(out)
		}
	}
	ratio := defBooleanTrueRatio
	if g.booleanTrueRatio != nil {
		ratio = *g.booleanTrueRatio
	}
	return g.chanceTrue(ratio)
}

func (g *generator) genUseNilPointer(t *Matcher) bool {
	for _, fn := range g.pointerNilFns {
		if out, ok := fn(t); ok {
			return g.chanceTrue(out)
		}
	}
	ratio := defNilPointerRatio
	if g.pointerNilRatio != nil {
		ratio = *g.pointerNilRatio
	}
	return g.chanceTrue(ratio)
}

func (g *generator) genString(t *Matcher) string {
	for _, fn := range g.stringFns {
		if out, ok := fn(t); ok {
			return out
		}
	}
	stringLen := g.genStringLen(t)
	if stringLen == 0 {
		return ""
	}
	source := getDefRunes()
	if len(g.runes) != 0 {
		source = g.runes
	}
	for _, fn := range g.runesFns {
		if out, ok := fn(t); ok {
			source = out
		}
	}
	return g.fillString(stringLen, source)
}

func (g *generator) fillString(size int, source []rune) string {
	runes := make([]rune, size)
	for j := 0; j < size; j++ {
		var i int
		if math.MaxInt == math.MaxInt32 {
			i = int(g.Uint32n(uint32(len(source))))
		} else {
			i = int(g.Uint64n(uint64(len(source))))
		}
		runes[j] = source[i]
	}
	return string(runes)
}

func genNumeric[T numeric](set nset[T], t *Matcher, g *generator) T {
	mm := defaultInterval[T]()
	if set.interval != nil {
		mm = *set.interval
	}
	for _, fn := range set.fns {
		if min, max, ok := fn(t); ok {
			mm.min = min
			mm.max = max
		}
	}
	if mm.min == mm.max {
		return mm.min
	}
	divisor := T(2)
	switch any(mm.min).(type) {
	case int, int64:
		return T(g.InclusiveInt64n(int64(mm.min), int64(mm.max)))
	case float32:
		return ((T(g.Float32()) * ((mm.max / divisor) - (mm.min / divisor))) + (mm.min / divisor)) * divisor
	case float64:
		return ((T(g.Float64()) * ((mm.max / divisor) - (mm.min / divisor))) + (mm.min / divisor)) * divisor
	}
	return T(g.InclusiveInt32n(int32(mm.min), int32(mm.max)))
}

func (g *generator) genFloat32(t *Matcher) float32 {
	return genNumeric(g.float32Set, t, g)
}

func (g *generator) genFloat64(t *Matcher) float64 {
	return genNumeric(g.float64Set, t, g)
}

func (g *generator) genStringLen(t *Matcher) int {
	return int(genNumeric(g.stringLenSet, t, g))
}

func (g *generator) genSliceLen(t *Matcher) int {
	return int(genNumeric(g.sliceLenSet, t, g))
}

func (g *generator) genMapLen(t *Matcher) int {
	return int(genNumeric(g.mapLenSet, t, g))
}

func (g *generator) genInt(t *Matcher) int {
	return genNumeric(g.intSet, t, g)
}

func (g *generator) genInt8(t *Matcher) int8 {
	return genNumeric(g.int8Set, t, g)
}

func (g *generator) genInt16(t *Matcher) int16 {
	return genNumeric(g.int16Set, t, g)
}

func (g *generator) genInt32(t *Matcher) int32 {
	return genNumeric(g.int32Set, t, g)
}

func (g *generator) genInt64(t *Matcher) int64 {
	return genNumeric(g.int64Set, t, g)
}

func (g *generator) genUint(t *Matcher) uint {
	return genNumeric(g.uintSet, t, g)
}

func (g *generator) genUint8(t *Matcher) uint8 {
	return genNumeric(g.uint8Set, t, g)
}

func (g *generator) genUint16(t *Matcher) uint16 {
	return genNumeric(g.uint16Set, t, g)
}

func (g *generator) genUint32(t *Matcher) uint32 {
	return genNumeric(g.uint32Set, t, g)
}

func (g *generator) genUint64(t *Matcher) uint64 {
	return genNumeric(g.uint64Set, t, g)
}
