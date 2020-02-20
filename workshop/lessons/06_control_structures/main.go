package main

import (
	"fmt"
)

func main() {

	// blocks

	outsideOfBlock := "from outer space"
	{
		insideOfBlock := "from inner space"

		fmt.Println(insideOfBlock)
		fmt.Println(outsideOfBlock)
	}
	// ⭣⭣⭣ out-comment + go run main.go ... what happens?
	// fmt.Println(insideOfBlock)

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
		case 2, 3:
			fmt.Println("Two AND Three, i is:", i)
		case 4:
			fmt.Println("Four")
			fallthrough
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	//labels

	//breake statement
	fmt.Println("")
	fmt.Println("breake statement:")
Loop:
	for i := 0; i < 100; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			if i == 2 {
				break Loop
			}
		}
	}

	//continue statement
	fmt.Println("")
	fmt.Println("continue statement:")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i=%v, j=%v\n", i, j)
			continue OuterLoop
		}
	}

	//goto statement
	fmt.Println("")
	fmt.Println("goto statement:")
	goto Done1
	fmt.Println("I am never reached")
Done1:
	fmt.Println("Done1")

	//impossible
	// 	goto Done2
	//     v := 0
	// Done2:
	//     fmt.Println(v)

	//impossible
	// goto Block
	// {
	// Block:
	// 	v := 0
	// 	fmt.Println(v)
	// }

	//impossible
	// goto Anonymous
	// func() {
	// Anonymous:
	// 	fmt.Println("anonymous func")
	// }()

	//impossible
	// 	func() {
	// 		fmt.Println("Nested function")
	// 		goto End
	// 	}()
	// End:
}
