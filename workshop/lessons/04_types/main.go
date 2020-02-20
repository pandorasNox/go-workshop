package main

import (
	"fmt"
)

func main() {
	//numbers
	fmt.Println("Numbers:")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)

	//strings
	fmt.Println("Strings:")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(`Hell World
		foo
	`)

	//booleans
	fmt.Println("Booleans:")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
