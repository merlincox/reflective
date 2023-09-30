package generator

import "reflect"

// Namable defines methods to enable customisation depending on type structures
type Namable interface {
	Name() string
	PkgPath() string
	Kind() reflect.Kind
}

type structField struct {
	field reflect.StructField
}

func (w structField) Kind() reflect.Kind {
	return reflect.Invalid
}

func (w structField) PkgPath() string {
	return w.field.PkgPath
}

func (w structField) Name() string {
	return w.field.Name
}

type indicator struct {
	value Indication
}

func (w indicator) Kind() reflect.Kind {
	return reflect.Invalid
}

func (w indicator) PkgPath() string {
	return ""
}

func (w indicator) Name() string {
	return string(w.value)
}

func newIndicator(value Indication) indicator {
	return indicator{value: value}
}

func pushNamed(visited []Namable, named Namable) []Namable {
	pushed := make([]Namable, len(visited)+1)
	copy(pushed, visited)
	pushed[len(visited)] = named
	return pushed
}
