package generator

import (
	"reflect"
)

type Matcher struct {
	parent     *Matcher
	current    reflect.Type
	field      *reflect.StructField
	sliceIndex int
	mapKey     bool
	mapValue   bool
	real       bool
	imaginary  bool
	mapLen     bool
	sliceLen   bool
}

func (t *Matcher) MatchesA(a any) bool {
	return t.current == reflect.TypeOf(a)
}

func (t *Matcher) MatchesAFieldOf(a any, name string) bool {
	if t.parent == nil || t.parent.Kind() != reflect.Struct {
		return false
	}
	if t.parent.current != reflect.TypeOf(a) {
		return false
	}
	if t.parent.field == nil || t.parent.field.Name != name {
		return false
	}

	return true
}

func (t *Matcher) IsAMapKeyOf(a any) bool {
	if t.parent == nil || t.parent.Kind() != reflect.Map {
		return false
	}
	if t.parent.current != reflect.TypeOf(a) {
		return false
	}

	return t.parent.mapKey
}

func (t *Matcher) IsAMapValueOf(a any) bool {
	if t.parent == nil || t.parent.Kind() != reflect.Map {
		return false
	}
	if t.parent.current != reflect.TypeOf(a) {
		return false
	}

	return t.parent.mapValue
}

func (t *Matcher) String() string {
	return t.current.String()
}

func (t *Matcher) Name() string {
	return t.current.Name()
}

func (t *Matcher) PkgPath() string {
	return t.current.PkgPath()
}

func (t *Matcher) Kind() reflect.Kind {
	return t.current.Kind()
}

func (t *Matcher) Parent() *Matcher {
	return t.parent
}

func (t *Matcher) HasParent() bool {
	return t.parent != nil
}

func (t *Matcher) forType(current reflect.Type) *Matcher {
	return &Matcher{
		current: current,
		parent:  t,
	}
}

func (t *Matcher) forField(current reflect.Type, field reflect.StructField) *Matcher {
	return &Matcher{
		current: current,
		field:   &field,
		parent:  t,
	}
}

func (t *Matcher) forMapKey(current reflect.Type) *Matcher {
	return &Matcher{
		current: current,
		mapKey:  true,
		parent:  t,
	}
}

func (t *Matcher) forMapValue(current reflect.Type) *Matcher {
	return &Matcher{
		current:  current,
		mapValue: true,
		parent:   t,
	}
}

func (t *Matcher) forSlice(current reflect.Type, sliceIndex int) *Matcher {
	return &Matcher{
		current:    current,
		sliceIndex: sliceIndex,
		parent:     t,
	}
}

func (t *Matcher) forReal(current reflect.Type) *Matcher {
	return &Matcher{
		current: current,
		real:    true,
		parent:  t,
	}
}

func (t *Matcher) forImaginary(current reflect.Type) *Matcher {
	return &Matcher{
		current:   current,
		imaginary: true,
		parent:    t,
	}
}

func (t *Matcher) forMapLen(current reflect.Type) *Matcher {
	return &Matcher{
		current: current,
		mapLen:  true,
		parent:  t,
	}
}

func (t *Matcher) forSliceLen(current reflect.Type) *Matcher {
	return &Matcher{
		current:  current,
		sliceLen: true,
		parent:   t,
	}
}
