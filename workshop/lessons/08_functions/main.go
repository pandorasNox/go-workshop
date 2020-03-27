package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}

// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	fmt.Println(total / float64(len(xs)))
// }

func main() {
	// argsExample(5, 4, 3, 324324)

	// shorthandArgDecEXample(1, 2, "foo")

	// fmt.Println(returnExample1())
	// r1 := returnExample1()
	// fmt.Println(r1)
	// fmt.Println()

	// fmt.Println(returnExample2())
	// r2a, r2b, r2c := returnExample2()
	// fmt.Println(r2a, r2b, r2c)
	// fmt.Println()

	// fmt.Println(returnExample3())

	//anonymous functions
	// myFunc := func(x func() int) {
	// 	fmt.Println("first")
	// 	x()
	// 	fmt.Println(x())
	// }
	// myFunc(func() int { fmt.Println("second"); return 5 })

	//self executing anonymous function
	// func() {
	// 	fmt.Println("self executing anonymous function")
	// 	func() {
	// 		fmt.Println("self executing anonymous function 2")
	// 	}()
	// }()

	//methods on structs
	// xs := []float64{98, 93, 77, 82, 83}
	// fmt.Println(average(xs))
	// p := point{x: 5}
	// p.doSmt(15)
	// fmt.Println(p)
}

//public - first letter uppercase
func ImPublic() {}

//private funcs
func imPrivate() {}

func argsExample(args ...int) {
	fmt.Println(args, reflect.TypeOf(args))
}

// func shorthandArgDecEXample(x int, y int, s string) {
func shorthandArgDecEXample(x, y int, s string) {
	fmt.Println(x, y, s)
}

func returnExample1() int {
	return 5
}

func returnExample2() (int, int, int) {
	return 5, 5, 5
}

//named returnes
func returnExample3() (x int, y int, z int) {
	// x = 1
	// y = 2
	// z = 3
	return
}

func average(xs []float64) float64 {
	panic("Not Implemented")
}

//methods on structs
type point struct {
	x int
	y int
}

func (p *point) doSmt(y int) {
	p.y = y
	fmt.Println(p)
}
