package generator

import (
	"reflect"
)

type Matcher struct {
	parent      *Matcher
	currentType reflect.Type
	field       *reflect.StructField
	sliceIndex  int
	mapKey      bool
	mapValue    bool
	real        bool
	imaginary   bool
	mapLen      bool
	sliceLen    bool
	name        string
}

func indirect(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

func matches(t1, t2 reflect.Type) bool {
	return indirect(t1) == indirect(t2)
}

func (t *Matcher) MatchesA(a any) bool {
	return matches(t.currentType, reflect.TypeOf(a))
}

func (t *Matcher) MatchesAFieldOf(a any, names ...string) bool {
	if t.parent == nil || t.parent.field == nil {
		return false
	}
	if !matches(t.parent.currentType, reflect.TypeOf(a)) {
		return false
	}
	for _, name := range names {
		if t.parent.field.Name == name {
			return true
		}
	}
	return false
}

func (t *Matcher) IsAMapKeyOf(a any) bool {
	if t.parent == nil || t.parent.Kind() != reflect.Map {
		return false
	}
	if !matches(t.parent.currentType, reflect.TypeOf(a)) {
		return false
	}
	return t.parent.mapKey
}

func (t *Matcher) IsAMapValueOf(a any) bool {
	if t.parent == nil || t.parent.Kind() != reflect.Map {
		return false
	}
	if !matches(t.parent.currentType, reflect.TypeOf(a)) {
		return false
	}
	return t.parent.mapValue
}

func (t *Matcher) String() string {
	return t.currentType.String()
}

func (t *Matcher) Name() string {
	return t.currentType.Name()
}

func (t *Matcher) PkgPath() string {
	return t.currentType.PkgPath()
}

func (t *Matcher) Kind() reflect.Kind {
	return indirect(t.currentType).Kind()
}

func (t *Matcher) Parent() *Matcher {
	return t.parent
}

func (t *Matcher) HasParent() bool {
	return t.parent != nil
}

func (t *Matcher) forType(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		name:        current.String(),
		parent:      t,
	}
}

func (t *Matcher) forField(current reflect.Type, field reflect.StructField) *Matcher {
	return &Matcher{
		currentType: current,
		field:       &field,
		parent:      t,
	}
}

func (t *Matcher) forMapKey(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		mapKey:      true,
		parent:      t,
	}
}

func (t *Matcher) forMapValue(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		mapValue:    true,
		parent:      t,
	}
}

func (t *Matcher) forSlice(current reflect.Type, sliceIndex int) *Matcher {
	return &Matcher{
		currentType: current,
		sliceIndex:  sliceIndex,
		parent:      t,
	}
}

func (t *Matcher) forReal(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		real:        true,
		parent:      t,
	}
}

func (t *Matcher) forImaginary(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		imaginary:   true,
		parent:      t,
	}
}

func (t *Matcher) forMapLen(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		mapLen:      true,
		parent:      t,
	}
}

func (t *Matcher) forSliceLen(current reflect.Type) *Matcher {
	return &Matcher{
		currentType: current,
		sliceLen:    true,
		parent:      t,
	}
}
