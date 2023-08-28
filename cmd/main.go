package main

import (
	"fmt"
	"os"

	"github.com/merlincox/reflective"
)

type TestStruct struct {
	Name    string
	Names   []string
	ID      int
	SomeMap map[int]string
	Fixed   [4]string
	Another *AnotherTestStruct
}

type AnotherTestStruct struct {
	Name string
	ID   int32
}

func main() {

	tester := TestStruct{}

	if err := reflective.FillRandomly(&tester); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", tester)
}
