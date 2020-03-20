package main

import (
	"fmt"
	"reflect"
)

func main() {
	//numbers
	fmt.Println("Numbers:")
	fmt.Println("1 + 1 =", 1+1, reflect.TypeOf(1+1))
	fmt.Println("1 + 1 =", 1.0+1, reflect.TypeOf(1.1+1))

	//strings
	fmt.Println("Strings:")
	fmt.Println(len("Hello World"))
	fmt.Println(len("ÜHello World"))
	fmt.Println(len([]rune("ÜHello World")))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(`Hell World
		foo
	`)

	//booleans
	fmt.Println("Booleans:")
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
