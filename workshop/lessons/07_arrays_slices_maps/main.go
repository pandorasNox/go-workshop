package main

import (
	"fmt"
	"reflect"
)

func main() {
	// arrays
	var foo [5]int
	fmt.Println("type is: " + reflect.TypeOf(foo).String())
	//printSlice(foo) //=> error, cannot use foo (type [5]int) as type []int
	fmt.Println("")

	// slices - make

	bar := make([]int, 3, 6)
	fmt.Println("type is: " + reflect.TypeOf(bar).String())
	printSlice(bar)
	fmt.Println("")

	bar = append(bar, 1, 2, 3)
	printSlice(bar)
	fmt.Println("")

	bar = append(bar, 1, 2, 3)
	printSlice(bar)
	fmt.Println("")

	// slices - assign

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// twoD := [][]int{[]int{}, []int{}}
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//maps

	// var x map[string]int // runtime error
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y[1])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
