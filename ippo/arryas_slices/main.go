package main

import (
	"fmt"
)

func main() {
	// var x [3]string = [3]string{"hello", "to", "world"}
	x := [3]string{"hello", "to", "world"}
	var y string = "hello"
	fmt.Println("eggs", "chocolate", "bananas")
	fmt.Println(x, y)

	x[2] = "africa"
	fmt.Println(x)
}
