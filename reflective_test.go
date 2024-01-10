package reflective_test

import (
	"github.com/merlincox/reflective/generator"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type Example struct {
	ScriptMap map[string]string
	Slice     []int
	Array     [10]int
	Complex1  complex64
	Complex2  complex128
	Unsafe    *unsafe.Pointer
}

func TestNonPointer(t *testing.T) {
	example := Example{}
	err := generator.New().Fill(example)
	assert.NotNil(t, err)
}
