package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//for's
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//if
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
			fmt.Println("divisible by 4:", i)
		}
	}

	//switch
	for i := 1; i <= 10; i++ {
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
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

}
