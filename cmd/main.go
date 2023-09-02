package main

import (
	"encoding/json"
	"fmt"
	"github.com/merlincox/reflective/generator"
	"os"
)

type TestStruct struct {
	Name     string             `json:"name"`
	Names    []string           `json:"names"`
	ID       int                `json:"id"`
	SomeMap  map[int]string     `json:"some_map"`
	Fixed    [4]string          `json:"fixed"`
	Another  *AnotherTestStruct `json:"another"`
	Another2 *AnotherTestStruct `json:"another2"`
	Bool     bool               `json:"bool"`
}

type AnotherTestStruct struct {
	Name string  `json:"name"`
	ID   int32   `json:"id"`
	Fval float32 `json:"fval"`
}

func main() {

	tester := TestStruct{}

	//if err := reflective.FillRandomly(&tester); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	//data, _ := json.MarshalIndent(tester, "", "  ")

	//fmt.Println(string(data))

	g := generator.New(
		generator.WithPointerNilChance(1),
		generator.WithBooleanFalseChance(1),
		generator.WithRunes([]rune("abc ")),
		generator.WithStringLenRange(0, 5),
		generator.WithSliceLenRange(0, 5),
		generator.WithMapLenRange(0, 5),
	)

	if err := g.FillRandomly(&tester); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, _ := json.MarshalIndent(tester, "", "  ")

	fmt.Println(string(data))

}
