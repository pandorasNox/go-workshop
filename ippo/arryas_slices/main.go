package main

import (
	"fmt"
)

//String is a slice of byte
func main() {
	x := "ippokratis"

	for i := 0; 1 < len(x); i++ {
		fmt.Print(x[i], " ")
	}

	fmt.Println()

}
