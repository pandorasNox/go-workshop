package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, playground")
	//fmt.Println(product(point{x:5, y:7}))
	//p := point{x:2, y:2}
	//fmt.Println(p.product())
	list := foo{1, 2, 3}
	fmt.Println("", list.sum())
}

type point struct {
	x int
	y int
}

type foo []int

func (f foo) sum() (result int) {
	sum := 0
	for _, v := range f {
		sum += v
		//sum = sum + v
	}

	return sum
}

//method on type point
func (p point) product() (result int) {
	return p.x + p.y
}

//a function which takes a point as an argument
func product(p point) (result int) {
	return p.x + p.y
}

//https://tour.golang.org/methods/1
