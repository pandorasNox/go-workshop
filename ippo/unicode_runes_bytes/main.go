package main
import (
	"fmt"
	"reflect"
)
func main() {
	//stringExampleOne()
	var sli := []int{1,2,3}
	for i,v := range sli {
	}
}
func stringExampleOne() {
	s := "hey"
	fmt.Println(s[0], reflect.TypeOf(s[0]))
	fmt.Println(s[1], reflect.TypeOf(s[1]))
	fmt.Println(s[2], reflect.TypeOf(s[2]))
	//fmt.Println(s[3], reflect.TypeOf(s[3]))
	fmt.Println()
	s1 := "Üey"
	fmt.Println(s1[0], reflect.TypeOf(s1[0]))
	fmt.Println(s1[1], reflect.TypeOf(s1[1]))
	fmt.Println(s1[2], reflect.TypeOf(s1[2]))
	fmt.Println(s1[3], reflect.TypeOf(s1[3]))
	fmt.Println()
	var b byte = 102
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println()
	sb := []byte{104, 101, 121}
	fmt.Println(sb, string(sb), reflect.TypeOf(sb))
	fmt.Println()
}