package generator

import (
	"fmt"
	"math"
	"reflect"
)

const (
	defMinStrLen        = 4
	defMaxStrLen        = 16
	defMinSliceLen      = 2
	defMaxSliceLen      = 16
	defMinMapLen        = 2
	defMaxMapLen        = 16
	defNilPointerRatio  = 0.5
	defBooleanTrueRatio = 0.5
	defMaxInt           = int(math.MaxInt8)
	defMaxFloat         = float64(math.MaxInt8)
)

func getDefRunes() []rune {
	return []rune("abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

type stringLenInt int
type mapLenInt int
type sliceLenInt int

type numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | stringLenInt | mapLenInt | sliceLenInt
}

type generator struct {
	rand Randomiser

	runes            []rune
	stringFns        []func(t *Matcher) (string, bool)
	booleanTrueRatio *float64
	boolTrueFns      []func(t *Matcher) (float64, bool)
	pointerNilRatio  *float64
	pointerNilFns    []func(t *Matcher) (float64, bool)
	runesFns         []func(t *Matcher) ([]rune, bool)

	stringLenSet nset[stringLenInt]
	mapLenSet    nset[mapLenInt]
	sliceLenSet  nset[sliceLenInt]
	float32Set   nset[float32]
	float64Set   nset[float64]
	intSet       nset[int]
	int8Set      nset[int8]
	int16Set     nset[int16]
	int32Set     nset[int32]
	int64Set     nset[int64]
	uintSet      nset[uint]
	uint8Set     nset[uint8]
	uint16Set    nset[uint16]
	uint32Set    nset[uint32]
	uint64Set    nset[uint64]
}

type nset[T numeric] struct {
	interval *interval[T]
	fns      []func(t *Matcher) (T, T, bool)
}

type interval[T numeric] struct {
	min T
	max T
}

// Option defines an option for customising the generator behaviour
type Option func(*generator) (*generator, error)

// New creates a new generator
func New() *generator {
	return new(generator)
}

// WithOptions adds options to a generator, returning the customised generator
func (g *generator) WithOptions(options ...Option) (*generator, error) {
	var err error
	for _, o := range options {
		g, err = o(g)
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

// Fill fills a data structure, by default pseudo-randomly. Its argument must be a pointer to the structure.
func (g *generator) Fill(a any) error {

	value := reflect.ValueOf(a)
	if value.Kind() != reflect.Pointer {
		return fmt.Errorf("the argument to Fill to must be a pointer")
	}

	g.fill(value.Elem(), nil)
	return nil
}

func (g *generator) fill(value reflect.Value, matcher *Matcher) {
	if !value.CanSet() {
		return
	}
	rtype := value.Type()

	switch value.Kind() {
	case reflect.Pointer:
		if g.genUseNilPointer(matcher.forSimpleType(rtype)) {
			return
		}
		value.Set(reflect.New(value.Type().Elem()))
		g.fill(value.Elem(), matcher.forSimpleType(rtype))

	case reflect.Bool:
		randBool := g.genBool(matcher.forSimpleType(rtype))
		value.SetBool(randBool)

	case reflect.Int:
		randInt := int64(g.genInt(matcher.forSimpleType(rtype)))
		value.SetInt(randInt)

	case reflect.Int8:
		randInt := int64(g.genInt8(matcher.forSimpleType(rtype)))
		value.SetInt(randInt)

	case reflect.Int16:
		randInt := int64(g.genInt16(matcher.forSimpleType(rtype)))
		value.SetInt(randInt)

	case reflect.Int32:
		randInt := int64(g.genInt32(matcher.forSimpleType(rtype)))
		value.SetInt(randInt)

	case reflect.Int64:
		randInt := g.genInt64(matcher.forSimpleType(rtype))
		value.SetInt(randInt)

	case reflect.Uint:
		randUint := uint64(g.genUint(matcher.forSimpleType(rtype)))
		value.SetUint(randUint)

	case reflect.Uint8:
		randUint := uint64(g.genUint8(matcher.forSimpleType(rtype)))
		value.SetUint(randUint)

	case reflect.Uint16:
		randUint := uint64(g.genUint16(matcher.forSimpleType(rtype)))
		value.SetUint(randUint)

	case reflect.Uint32:
		randUint := uint64(g.genUint32(matcher.forSimpleType(rtype)))
		value.SetUint(randUint)

	case reflect.Uint64:
		randUint := g.genUint64(matcher.forSimpleType(rtype))
		value.SetUint(randUint)

	case reflect.Float32:
		randFloat := float64(g.genFloat32(matcher.forSimpleType(rtype)))
		value.SetFloat(randFloat)

	case reflect.Float64:
		randFloat := g.genFloat64(matcher.forSimpleType(rtype))
		value.SetFloat(randFloat)

	case reflect.Complex64:
		r := g.genFloat32(matcher.forRealPart(rtype))
		i := g.genFloat32(matcher.forImaginaryPart(rtype))
		value.SetComplex(complex128(complex(r, i)))

	case reflect.Complex128:
		r := g.genFloat64(matcher.forRealPart(rtype))
		i := g.genFloat64(matcher.forImaginaryPart(rtype))
		value.SetComplex(complex(r, i))

	case reflect.String:
		randStringVal := g.genString(matcher.forSimpleType(rtype))
		value.SetString(randStringVal)

	case reflect.Slice:
		elementType := rtype.Elem()
		size := g.genSliceLen(matcher.forSliceLen(rtype))
		sliceVal := reflect.MakeSlice(reflect.SliceOf(elementType), 0, size)
		for i := 0; i < size; i++ {
			newElement := reflect.Indirect(reflect.New(elementType))
			g.fill(newElement, matcher.forSliceElement(rtype, i, size))
			sliceVal = reflect.Append(sliceVal, newElement)
		}
		value.Set(sliceVal)

	case reflect.Array:
		for i := 0; i < value.Len(); i++ {
			g.fill(value.Index(i), matcher.forArrayElement(rtype, i, value.Len()))
		}

	case reflect.Map:
		mapVal := reflect.MakeMap(rtype)
		size := g.genMapLen(matcher.forMapLen(rtype))
		// note that actual map length will be lower than size if any duplicate keys are generated
		for i := 0; i < size; i++ {
			newKey := reflect.Indirect(reflect.New(rtype.Key()))
			g.fill(newKey, matcher.forMapKey(rtype))
			newElement := reflect.Indirect(reflect.New(rtype.Elem()))
			g.fill(newElement, matcher.forMapElement(rtype, newKey.Interface()))
			mapVal.SetMapIndex(newKey, newElement)
		}
		value.Set(mapVal)

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			g.fill(value.Field(i), matcher.forField(rtype, rtype.Field(i)))
		}
	}
}
