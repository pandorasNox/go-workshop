package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// for's
	for i := 1.0; i <= 500; i += 0.5 {
		// i = i + 0.5
		fmt.Println(i)
	}

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("index:%d | value:%d\n", i, v)
		fmt.Println("index:", i, " | value:\n", v)
	}

	// if
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even:", i)
		} else {
			fmt.Println("odd", i)
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("divisible by 2:", i)
		} else if i%3 == 0 {
			fmt.Println("divisible by 3:", i)
		} else if i%4 == 0 {
			//dead code example, not captured by compiler
			fmt.Println("divisible by 4:", i)
		} else {
			fmt.Println("unkown")
		}
	}

	//switch
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
			fallthrough
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

}
