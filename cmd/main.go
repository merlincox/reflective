package main

import (
	"encoding/json"
	"fmt"
	"github.com/merlincox/reflective/generator"
	"os"
)

type SomeEnum int

const (
	enum1 SomeEnum = iota
	enum2
	enum3
)

type FuncType func()

type TestStruct struct {
	Name        string             `json:"name"`
	Names       []string           `json:"names"`
	ID          int                `json:"id"`
	SomeMap     map[int]string     `json:"some_map"`
	Fixed       [4]string          `json:"fixed"`
	Another     *AnotherTestStruct `json:"another"`
	Another2    *AnotherTestStruct `json:"another2"`
	Bool        bool               `json:"bool"`
	Fval2       float32            `json:"fval2"`
	Enumval     SomeEnum           `json:"enumval"`
	SomeFunc    func()             `json:"-"`
	AnotherFunc *FuncType          `json:"-"`
}

type AnotherTestStruct struct {
	Name  string  `json:"name"`
	ID    int32   `json:"id"`
	Fval  float32 `json:"fval"`
	Fval2 float32 `json:"fval2"`
}

func main() {

	tester := TestStruct{}

	g := generator.New(
		generator.PointerNilChance(0),
		generator.BooleanFalseChance(1),
		generator.Runes([]rune("abc ")),
		generator.StringLenRange(1, 16),
		generator.SliceLenRange(1, 5),
		generator.MapLenRange(1, 5),
		generator.Float32Fn(
			func(visited ...generator.Namable) (float32, bool) {
				l := len(visited)
				if l > 2 {
					if visited[l-3].Name() == "AnotherTestStruct" && visited[l-2].Name() == "Fval2" {
						return float32(0.7), true
					}
				}
				return 0, false
			}),
	)

	g = g.WithOption(
		generator.IntFn(
			func(visited ...generator.Namable) (int, bool) {
				l := len(visited)
				if l > 0 {
					if visited[l-1].PkgPath() == "main" && visited[l-1].Name() == "SomeEnum" {
						return g.Intn(int(enum3) + 1), true
					}
				}
				return 0, false
			}))

	if err := g.FillRandomly(&tester); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", tester)
	data, _ := json.MarshalIndent(tester, "", "  ")

	fmt.Println(string(data))
}
