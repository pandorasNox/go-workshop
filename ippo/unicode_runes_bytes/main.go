package main

import (
	"fmt"
	"reflect"
)

func main() {
	//stringExampleOne()
	runeExample()
}

func runeExample() {
	// fmt.Println("Hello")
	// str := "Üey"
	//fmt.Println(str[0], str[1], "'"+string(str[0])+"'")

	//rune is an alias for int32
	//casting rune (casting in32)
	//into a string => gives bac the unicode char/string
	//rune represents the start int number of the unicode character
	// ru := rune(177)
	// fmt.Println("rune:", ru, reflect.TypeOf(ru))
	// ru2 := 'Ü'
	// fmt.Println("rune:", ru2, string(ru2), reflect.TypeOf(ru2))

	str2 := "Üey1"
	for i, singleRune := range str2 {
		if singleRune == '1' {
			fmt.Println("HIT")
		}
		fmt.Printf("index: %d | rune: %v,%s\n", i, singleRune, string(singleRune))
	}

	// s := "abc" // => a - 0 ; b - 1; c - 2
	// slii := []int{1,5,8} // 1 - 0; 5 - 1; 8 - 2
	// // for _,_ := range s {
	// for i,v := range slii {
	// 	//execute code
	// }
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
