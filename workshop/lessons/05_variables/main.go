package main

import "fmt"

const (
	d    int8 = 5
	e         = 10
	f         = 15
	bars      = "string"
)

func main() {

	// initialises it with the zero value of the variable's type
	var x int
	fmt.Println("x", x)

	var y int
	y = 5
	fmt.Println("y:", y)

	var w = 9
	fmt.Println("w:", w)

	//common
	var u int8 = 7
	fmt.Println("7:", u)

	//common
	z := 3
	fmt.Println("z:", z)

	//constants
	const t string = "Hello World"
	fmt.Println("t:", t)
	// t = "Some other string" // check what happens when you outcomment this and compile

	//shorthand
	var (
		a   int8 = 5
		b        = 10
		c        = 15
		foo      = "string"
	)
	fmt.Println("a, b, c:", a, b, c, foo)
}
