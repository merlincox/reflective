package main

import (
	"encoding/json"
	"fmt"
	"github.com/merlincox/reflective/generator"
	"os"
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
	Field1 int
	Field2 int
}

func main() {
	tester := TestStruct{}

	g := generator.New(
		generator.IntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.MatchesAFieldOf(SubStruct{}, "Field1") {
					return -99, -99, true
				}
				return 0, 0, false
			}),
	)
	g = g.WithOptions(

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

	data, _ := json.MarshalIndent(tester, "", "  ")

	fmt.Println(string(data))
}
