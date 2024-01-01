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
	//Enumval      SomeEnum
	//SubStructVal SubStruct
	//MapVal       map[string]int
	SliceVal  []*SubStruct
	SliceVal2 []*AnonStruct
	//Unknown      any
	//Unknown2     any
	//Unknown3     any
	//AnonStruct
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
	g := generator.New()

	target := g.Float64()
	fmt.Printf("target: %v %%\n", target*100)
	max := 1000000000
	hits := 0

	for i := 1; i <= max; i++ {
		hit := g.Float64() <= target
		if hit {
			hits++
		}
	}

	actual := float64(hits) / float64(max)
	fmt.Printf("actual: %v %%\n", actual*100)
	fmt.Printf("divergence: %v %%\n", (target-actual)*100)

}

func main2() {
	tester := TestStruct{}

	g := generator.New()

	g, _ = g.WithOptions(
		generator.WithIntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.MatchesAFieldOf(&SubStruct{}, "Field1", "Field2") {
					return -99, -99, true
				}
				return 0, 0, false
			}),
	)
	g, _ = g.WithOptions(

		generator.WithPointerNilRatioFn(
			func(t *generator.Matcher) (float64, bool) {
				if t.MatchesA(&SubStruct{}) {
					return 0, true
				}
				return 1, true
			},
		),

		generator.WithIntFn(
			func(m *generator.Matcher) (int, int, bool) {
				if m.MatchesA(enumMin) {
					return int(enumMin), int(enumMax), true
				}
				return 0, 0, false
			}),

		generator.WithIntFn(
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
