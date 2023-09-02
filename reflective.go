package reflective

import (
	"github.com/merlincox/reflective/generator"
)

func FillRandomly(a any) error {
	c := generator.New()
	return c.FillRandomly(a)
}
