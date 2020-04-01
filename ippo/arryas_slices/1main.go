package main

import (
	"fmt"
)

func main() {
	// var x [3]string = [3]string{"hello", "to", "world"}
	//s := [3]string{"hello", "to", "world"}
	//var y string = "hello"
	//	fmt.Println("eggs", "chocolate", "bananas")
	//	fmt.Println(s, y)
	//
	//	s[2] = "africa"
	//	fmt.Println(s)
	//	intList()
	//	fmt.Println(fetchList())
	//
	//	intSlice()
	//	fmt.Println(fetchSlice())
	fmt.Println(fetchSlice1())

}

func intList() {
	z := [5]int{1, 2, 3, 4, 5}
	fmt.Println(z)

}

func fetchList() [5]int {
	z := [5]int{5, 4, 3, 2, 1}
	return z
}

func intSlice() {
	z := []int{6, 7, 8, 9, 10}
	fmt.Println(z)

}

func fetchSlice() []int {
	z := []int{10, 9, 8, 7, 6}
	return z

}

func fetchSlice1() []int {
	//var s0 []int
	//printSlice(s0)
	//s0 = append(s0, 1)
	//printSlice(s0)
	//s0 = append(s0, 1)
	//printSlice(s0)
	//s0 = append(s0, 1)
	//printSlice(s0)
	//s0 = append(s0, 1)
	//printSlice(s0)
	//s0 = append(s0, 1)
	//printSlice(s0)

	//s1 := []int{}
	//printSlice(s1)
	//fmt.Println()

	//s2 := make([]int, 0)
	//printSlice(s2)
	//fmt.Println()

	//s3 := make([]int, 4)
	//printSlice(s3)
	//fmt.Println()

	fmt.Println("s4:")
	s4 := make([]int, 4)
	printSlice(s4)
	s4 = append(s4, 5)
	printSlice(s4)
	fmt.Println()

	fmt.Println("s5:")
	s5 := make([]int, 4, 10)
	s5[0] = 100
	printSlice(s5)
	s5 = append(s5, 5, 6, 7, 8)
	printSlice(s5)
	fmt.Println()

	fmt.Println("s6:")
	s6 := []int{0, 5, 8, 10, 15}
	printSlice(s6)
	s6 = append(s6, 5)
	printSlice(s6)
	fmt.Println()

	s := make([]int, 1, 10)
	s = append(s, 1, 3, 5)
	return s

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
