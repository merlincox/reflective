package reflective

import (
	"github.com/merlincox/reflective/generator"
)

// FillRandomly fills a data structure pseudo-randomly using default settings. The argument must be a pointer to the structure.
func FillRandomly(a any) error {
	return generator.New().Fill(a)
}
