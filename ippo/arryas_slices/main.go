package main

import (
	"fmt"
)

func main() {
	// var x [3]string = [3]string{"hello", "to", "world"}
	s := [3]string{"hello", "to", "world"}
	var y string = "hello"
	fmt.Println("eggs", "chocolate", "bananas")
	fmt.Println(s, y)

	s[2] = "africa"
	fmt.Println(s)
	intList()
	fmt.Println(fetchList())

}

func intList() {
	z := [5]int{1, 2, 3, 4, 5}
	fmt.Println(z)

}

func fetchList() [5]int {
	z := [5]int{5, 4, 3, 2, 1}
	return z
}
