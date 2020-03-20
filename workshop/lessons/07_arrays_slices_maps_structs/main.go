package main

import (
	"fmt"
	"reflect"
)

//Vertex ...
type Vertex struct {
	X int
	Y int
}

func main() {
	// arrays
	var foo [5]int
	fmt.Println("type is: " + reflect.TypeOf(foo).String())
	// printSlice(foo) //=> error, cannot use foo (type [5]int) as type []int
	fmt.Println("")

	// slices - make

	aSlice := make([]int, 3, 6)
	fmt.Println("type is: " + reflect.TypeOf(aSlice).String())
	printSlice(aSlice)
	fmt.Println("")

	aSlice = append(aSlice, 1, 2, 3)
	printSlice(aSlice)
	fmt.Println("")

	aSlice = append(aSlice, 1, 2, 3)
	printSlice(aSlice)
	fmt.Println("")

	// slices - assign

	t := []string{"g", "h", "i", "hello"}
	fmt.Println("dcl:", t)
	fmt.Println("type is: " + reflect.TypeOf(t).String())
	t = append(t, "world")
	fmt.Println("dcl:", t)
	fmt.Println(t[4])

	twoD := [][]int{[]int{}, []int{}, []int{}}
	// twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	fmt.Println("")

	//slicing

	// A slice can also be formed by "slicing" an existing slice or array.
	// Slicing is done by specifying a half-open range
	// with two indices separated by a colon. For example, the expression b[1:4]
	// creates a slice including elements 1 through 3 of b
	// (the indices of the resulting slice will be 0 through 2).
	//
	// The start and end indices of a slice expression are optional;
	// they default to zero and the slice's length
	fmt.Println("slicing:")
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	c := b[1:4]
	fmt.Println("b: ", b, &b)
	fmt.Println("b[1:4]: ", c, &c)
	fmt.Println("b[:]: ", b[:])
	fmt.Println("b[:4]: ", b[:4])
	fmt.Println("b[2:]: ", b[2:])
	fmt.Println("b[:6]: ", b[:6]) //no error, although index 6 doesn't exist
	// fmt.Println("b[:4]: ", b[:7]) //runtime error, slice bounds out of range [:7] with capacity 6
	fmt.Println("")

	//maps

	// var x map[string]int // runtime error
	x := make(map[string]int)
	x["age"] = 10
	fmt.Println(x)
	x["days"] = 200
	fmt.Println(x)

	y := make(map[int]int)
	y[1] = 10
	fmt.Println(y)
	fmt.Println(y[1])
	y[156] = 79869
	fmt.Println(y)
	fmt.Println(y[156])

	// structs

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p := &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, v2, v3, p)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
