package main

import (
	"fmt"
	"reflect"
	//"reflect"
)

func main() {
	fmt.Println("intro:")
	fmt.Printf("string ÜH[0]: %v\n", "ÜH"[0])
	fmt.Printf("string ÜH[1]: %v\n", "ÜH"[1])
	fmt.Printf("string ÜH[2]%v\n", "ÜH"[2])
	fmt.Printf("rune Ü: %v\n", 'Ü')
	fmt.Printf("rune H: %v\n", 'H')
	fmt.Printf("%v\n", 'Ü' == "Üh"[0])
	fmt.Println("rune (is an alias) type is: " + reflect.TypeOf('H').String())
	fmt.Println("byte (is an alias) type is: " + reflect.TypeOf("H"[0]).String())
	fmt.Println("rune (is an alias) type is: " + reflect.TypeOf(rune("H"[0])).String())
	fmt.Println("byte (is an alias) type is: " + reflect.TypeOf(byte("H"[0])).String())
	fmt.Println("")

	// rune slice

	var s string

	fmt.Println("runes:")
	s = string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -1})
	fmt.Println("string from rune slice:", s) // ABC€�
	s = string([]rune{'A', 'B', 'C', '€', '�'})
	fmt.Println("string from rune slice:", s) // ABC€�
	fmt.Printf("%U\n", []rune(s))
	fmt.Println("")

	//convertion
	fmt.Println("conversion:")

	// "for range" returns rune's, and the index where the rune starts
	fmt.Println("for:")
	for i, v := range "ÜH" {
		fmt.Printf("index: %v | value: %v | type: %v\n", i, v, reflect.TypeOf(v).String())
	}
	fmt.Println("")

	//string to rune slice
	s = "ÜH"
	sRunes := []rune(s)
	fmt.Println("rune slice: ", sRunes)
	fmt.Printf("%U\n", sRunes) // [U+0041 U+0042 U+0043 U+20AC]

	//rune slice to string:
	s = string(sRunes)
	fmt.Println("string from rune slice: ", s)

	//rune to string
	s = string('Ü')
	fmt.Println("string from rune: ", s)
	fmt.Println("len string from rune: ", len(s))

	//string to rune => not possible => can only be converted to "slice of rune's"
	// ru := rune(s)
}
