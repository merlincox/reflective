package generator

import (
	"reflect"
)

type Matcher struct {
	parent          *Matcher
	rtype           reflect.Type
	field           *reflect.StructField
	index           int
	isMapKey        bool
	isMapElement    bool
	mapKeyValue     any
	isRealPart      bool
	isImaginaryPart bool
	isSliceElement  bool
	isArrayElement  bool
	isMapLen        bool
	isSliceLen      bool
	name            string
	length          int
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
	return matches(t.rtype, reflect.TypeOf(a))
}

func (t *Matcher) MatchesAFieldOf(a any, names ...string) bool {
	if t.parent == nil || t.parent.field == nil {
		return false
	}
	if !matches(t.parent.rtype, reflect.TypeOf(a)) {
		return false
	}
	for _, name := range names {
		if t.parent.field.Name == name {
			return true
		}
	}
	return false
}

func (t *Matcher) IsAMapKey() bool {
	return t.parent != nil && t.parent.isMapKey
}

func (t *Matcher) IsAMapElement() bool {
	return t.parent != nil && t.parent.isMapElement
}

func (t *Matcher) IsASliceElement() bool {
	return t.parent != nil && t.parent.isSliceElement
}

func (t *Matcher) IsAnArrayElement() bool {
	return t.parent != nil && t.parent.isArrayElement
}

func (t *Matcher) IsARealPart() bool {
	return t.isRealPart
}

func (t *Matcher) IsAnImaginaryPart() bool {
	return t.isImaginaryPart
}

func (t *Matcher) Index() int {
	return t.index
}

func (t *Matcher) Length() int {
	return t.length
}

func (t *Matcher) MapKeyValue() any {
	return t.mapKeyValue
}

func (t *Matcher) Type() reflect.Type {
	return t.rtype
}

func (t *Matcher) Parent() *Matcher {
	return t.parent
}

func (t *Matcher) HasParent() bool {
	return t.parent != nil
}

func (t *Matcher) forSimpleType(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:  current,
		name:   current.String(),
		parent: t,
	}
}

func (t *Matcher) forField(current reflect.Type, field reflect.StructField) *Matcher {
	return &Matcher{
		rtype:  current,
		field:  &field,
		parent: t,
	}
}

func (t *Matcher) forMapKey(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:    current,
		isMapKey: true,
		parent:   t,
	}
}

func (t *Matcher) forMapElement(current reflect.Type, key any) *Matcher {
	return &Matcher{
		rtype:        current,
		isMapElement: true,
		mapKeyValue:  key,
		parent:       t,
	}
}

func (t *Matcher) forSliceElement(current reflect.Type, index int, length int) *Matcher {
	return &Matcher{
		rtype:          current,
		isSliceElement: true,
		index:          index,
		length:         length,
		parent:         t,
	}
}

func (t *Matcher) forArrayElement(current reflect.Type, index int, length int) *Matcher {
	return &Matcher{
		rtype:          current,
		isArrayElement: true,
		index:          index,
		length:         length,
		parent:         t,
	}
}

func (t *Matcher) forRealPart(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:      current,
		isRealPart: true,
		parent:     t,
	}
}

func (t *Matcher) forImaginaryPart(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:           current,
		isImaginaryPart: true,
		parent:          t,
	}
}

func (t *Matcher) forMapLen(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:    current,
		isMapLen: true,
		parent:   t,
	}
}

func (t *Matcher) forSliceLen(current reflect.Type) *Matcher {
	return &Matcher{
		rtype:      current,
		isSliceLen: true,
		parent:     t,
	}
}
