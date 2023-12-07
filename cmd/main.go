package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/merlincox/reflective/generator"
)

type SomeEnum int

const (
	enumMin SomeEnum = iota
	enum2
	enumMax
)

type TestStruct struct {
	Enumval      SomeEnum
	SubStructVal SubStruct
	MapVal       map[string]int
	SliceVal     []string
	Unknown      any
	Unknown2     any
	Unknown3     any
	AnonStruct
}

type SubStruct struct {
	Field1 int
	Field2 int
}

type AnonStruct struct {
	Field3 int
	Field4 int
}

func main() {
	tester := TestStruct{}

	g, _ := generator.New(
		generator.IntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.MatchesAFieldOf(&SubStruct{}, "Field1", "Field2") {
					return -99, -99, true
				}
				return 0, 0, false
			}),
	)
	g, _ = g.WithOptions(

		generator.IntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.MatchesA(enumMin) {
					return int(enumMin), int(enumMax), true
				}
				return 0, 0, false
			}),

		generator.IntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.IsAMapValueOf(map[string]int{}) {

					return -6, 6, true
				}
				return 0, 0, false
			}),
	)

	if err := g.FillRandomly(&tester); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(tester, "", "  ")
	if err != nil {
		fmt.Printf("=====\n%s\n", err.Error())
		os.Exit(1)
	}
	if len(data) == 0 {
		fmt.Println("empty")
	}
	fmt.Println(string(data))
}
