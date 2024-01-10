package generator_test

import (
	"github.com/merlincox/reflective/generator"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"unsafe"
)

const (
	latin         = "abcdefghijklmnoqrstuvwxyz"
	cyrillic      = "АБВГДЕЖꙂꙀИІКЛМНОПРСТОУФХѠЦЧШЩЪЪІЬѢꙖѤЮѪѬѦѨѮѰѲѴҀ"
	greek         = "ΑαΒβΓγΔδΕεΖζΗηΘθΙιΚκΛλΜμΝνΞξΟοΠπΡρΣσςΤτΥυΦφΧχΨψΩω"
	scriptsConcat = "Latin,Greek,Cyrillic"
)

type Example struct {
	ScriptMap map[string]string
	Slice     []int
	Array     [10]int
	Complex1  complex64
	Complex2  complex128
	Unsafe    *unsafe.Pointer
}

var scripts = strings.Split(scriptsConcat, ",")

func charsetFromKey(key string) (string, bool) {
	switch key {
	case "Latin":
		return latin, true
	case "Cyrillic":
		return cyrillic, true
	case "Greek":
		return greek, true
	}
	return "", false
}

func TestMapMatching(t *testing.T) {

	subject := generator.New()

	subject, _ = subject.WithOptions(

		generator.WithStringFn(
			func(t *generator.Matcher) (string, bool) {
				if t.IsAMapKey() {
					return scripts[int(subject.Uint32n(uint32(len(scripts))))], true
				}
				return "", false
			}),
		generator.WithRunesFn(
			func(t *generator.Matcher) ([]rune, bool) {
				if t.IsAMapElement() {
					key := t.Parent().MapKeyValue().(string)
					chars, ok := charsetFromKey(key)
					if ok {
						return []rune(chars), true
					}
				}
				return nil, false
			}),
	)
	example := new(Example)
	subject.Fill(example)
	found := false
	var value string
	for key := range example.ScriptMap {
		for _, script := range scripts {
			if key == script {
				found = true
				value = script
				break
			}
		}
	}
	assert.True(t, found)
	charset, ok := charsetFromKey(value)
	assert.True(t, ok)
	assert.True(t, strings.Contains(charset, example.ScriptMap[value][0:1]))
}

func TestSliceMatching(t *testing.T) {

	subject := generator.New()

	subject, _ = subject.WithOptions(
		generator.WithIntFn(
			func(t *generator.Matcher) (int, int, bool) {
				if t.IsASliceElement() {
					i := t.Parent().Index()
					return i, i, true
				}
				return 0, 0, false
			}),
	)
	example := new(Example)
	subject.Fill(example)
	for i, val := range example.Slice {
		assert.Equal(t, i, val)
	}
}

func TestArrayMatching(t *testing.T) {

	subject := generator.New()

	subject, _ = subject.WithOptions(
		generator.WithIntFn(
			func(t *generator.Matcher) (int, int, bool) {
				if t.IsAnArrayElement() {
					i := t.Parent().Length() - t.Parent().Index()
					return i, i, true
				}
				return 0, 0, false
			}),
	)
	example := new(Example)
	subject.Fill(example)
	for i, val := range example.Array {
		assert.Equal(t, 10-i, val)
	}
}

func TestComplexMatching(t *testing.T) {
	real1 := float32(99.9)
	imaginary1 := float32(44.9)
	real2 := 99.9
	imaginary2 := 44.9
	subject := generator.New()

	subject, _ = subject.WithOptions(
		generator.WithFloat64Fn(
			func(t *generator.Matcher) (float64, float64, bool) {
				if t.IsARealPart() {
					return real2, real2, true
				}
				if t.IsAnImaginaryPart() {
					return imaginary2, imaginary2, true
				}
				return 0, 0, false
			}),
		generator.WithFloat32Fn(
			func(t *generator.Matcher) (float32, float32, bool) {
				if t.IsARealPart() {
					return real1, real1, true
				}
				if t.IsAnImaginaryPart() {
					return imaginary1, imaginary1, true
				}
				return 0, 0, false
			}),
	)
	example := new(Example)
	subject.Fill(example)
	assert.Equal(t, real1, real(example.Complex1))
	assert.Equal(t, imaginary1, imag(example.Complex1))
	assert.Equal(t, real2, real(example.Complex2))
	assert.Equal(t, imaginary2, imag(example.Complex2))

}

func TestNonPointer(t *testing.T) {
	example := Example{}
	err := generator.New().Fill(example)
	assert.NotNil(t, err)
}
