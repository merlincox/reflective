package main

import (
	"fmt"
	"github.com/merlincox/reflective"
	"golang.org/x/exp/constraints"
	"os"
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

type Number interface {
	constraints.Integer | constraints.Float
}

func main() {

	//var n int64 = 2
	//p := 1
	//for i := 0; i <= 63; i++ {
	//	n = n * 2
	//	p++
	//	if p > 60 {
	//		fmt.Println(p, n)
	//	}
	//}
	//
	//fmt.Println(p, n)
	//res := math.Pow(2, 63)
	//fmt.Println(res)
	//var y int64
	//y = math.MaxInt64
	//fmt.Printf("MaxInt64 = %d\n", y)
	//fmt.Printf("MaxInt64 + 1 = %d\n", y+1)

	//var y int64 = math.MinInt64
	//fmt.Printf("math.MinInt64 = %d\n", y)
	//fmt.Printf("Diff(0, 5) = %d\n", Diff(0, 5))
	//fmt.Printf("Diff(5, 10) = %d\n", Diff(5, 10))
	//fmt.Printf("Diff(-5, 5) = %d\n", Diff(-5, 5))
	//fmt.Printf("Diff(-5, 0) = %d\n", Diff(-5, 0))
	//fmt.Printf("Diff(math.MinInt64, math.MaxInt64) = %d\n", Diff(math.MinInt64, math.MaxInt64))
	//fmt.Printf("Diff(math.MinInt64, 0) = %d\n", Diff(math.MinInt64, 0))
	//fmt.Printf("Diff(math.MinInt64, -1) = %d\n", Diff(math.MinInt64, -1))

	//var y int64 = math.MinInt64
	//fmt.Printf("y = %d\n", y)
	//fmt.Printf("1-y = %d\n", 1-y)
	//fmt.Printf("-(1-y) = %d\n", -(1 - y))
	//fmt.Printf("size = %d\n", uint64(-(1-y))+1)
	//y = -10
	//fmt.Printf("y = %d\n", y)
	//fmt.Printf("-y = %d\n", -y)
	//fmt.Printf("size = %d\n", uint64(-y))
	//
	//y = math.MinInt64 + 1
	//fmt.Printf("y = %d\n", y)
	//fmt.Printf("-y = %d\n", -y)
	//fmt.Printf("size = %d\n", uint64(-y))
	//
	//y = math.MinInt64
	//fmt.Printf("y = %d\n", y)
	//fmt.Printf("-y = %d\n", -y)
	//fmt.Printf("size = %d\n", uint64(-y))

	//fmt.Printf("uint64(MinInt64) = %d\n", uint64(y))
	//fmt.Printf("uint64(MinInt64 +1) = %d\n", uint64(y+1))
	//fmt.Printf("uint64(MinInt64 +100) = %d\n", uint64(y+100))
	//var zz int64 = -1
	//fmt.Printf("uint64(-1) = %d\n", uint64(zz))
	//var min, max int64 = -5, 10

	//diff := uint64(min) + uint64(max)
	//fmt.Printf("uint64(min) = %d\n", uint64(min))
	//fmt.Printf("uint64(max) = %d\n", uint64(max))
	//fmt.Printf("diff = %d\n", diff)
	//
	//var z uint64 = math.MaxUint64
	//fmt.Printf("MaxUint64 = %d\n", z)
	//
	//fmt.Printf("MaxUint64/2 = %d\n", z/2)
	//fmt.Printf("MaxUint64 - (MaxUint64/2) = %d\n", z-(z/2))
	//
	//fmt.Printf("MaxUint64-1/2 = %d\n", z/2)
	//fmt.Printf("MaxUint64-1 - (MaxUint64-1/2) = %d\n", z-(z/2))
	//
	//var hh int64 = math.MaxInt64
	//fmt.Printf("Maxint64 = %d\n", hh)
	//u := reflective.MapItou(hh)
	//fmt.Printf("Maxint64 as uint64 = %d\n", u)
	//
	//hh = math.MinInt64
	//fmt.Printf("Mixint64 = %d\n", hh)
	//hh = math.MinInt64
	//fmt.Printf("Mixint64>>2 = %d\n", hh>>2)
	//y := math.MinInt64
	//fmt.Printf("MinInt64 = %d\n", y)
	//y = math.MaxInt64
	//fmt.Printf("MaxInt64 = %d\n", y)
	//var z uint64 = math.MaxUint64
	//fmt.Printf("MaxUint64 = %d\n", z)
	//
	//var uu uint64 = 0
	//fmt.Printf("reflective.MapUtoi(0) = %d\n", reflective.MapUtoi(uu))
	//uu = 1
	//fmt.Printf("reflective.MapUtoi(1) = %d\n", reflective.MapUtoi(uu))
	//uu = math.MaxUint64
	//fmt.Printf("reflective.MapUtoi(math.MaxUint64) = %d\n", reflective.MapUtoi(uu))
	//
	//var xx int64 = math.MinInt64
	//fmt.Printf("reflective.MapItou(math.MinInt64)) = %d\n", reflective.MapItou(xx))
	//xx++
	//fmt.Printf("reflective.MapItou(math.MinInt64)++) = %d\n", reflective.MapItou(xx))
	//xx = math.MaxInt64
	//fmt.Printf("reflective.MapItou(math.MaxInt64)++) = %d\n", reflective.MapItou(xx))

	//fmt.Printf("MinInt64 + 1 = %d\n", y+1)
	//fmt.Printf("-(MinInt64 + 1) = %d\n", -(y + 1))

	//out := strings.Repeat("1", 63)
	//iout, err := strconv.ParseInt(out, 2, 64)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("iout = %d\n", iout)
	//}
	//out = strings.Repeat("1", 64)
	//iout, err = strconv.ParseInt(out, 2, 64)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("iout = %d\n", iout)
	//}
	//uout, err := strconv.ParseUint(out, 2, 64)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("uout = %d\n", uout)
	//}

	tester := TestStruct{}

	if err := reflective.FillRandomly(&tester); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", tester)
}
