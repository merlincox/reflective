package reflective

import (
	"reflect"

	"github.com/merlincox/reflective/generator"
)

// FillRandomly fills a data structure pseudo-randomly using default settings. The argument must be a pointer to the structure.
func FillRandomly(a any) error {
	c := generator.New()
	return c.FillRandomly(a)
}

// FillRandomlyByValue fills a data structure pseudo-randomly using default settings. The argument must be the reflect.Vaue of a pointer to the structure.
func FillRandomlyByValue(val reflect.Value) error {
	c := generator.New()
	return c.FillRandomlyByValue(val)
}
