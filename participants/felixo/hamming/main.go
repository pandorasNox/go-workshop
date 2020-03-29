package main

import (
	"fmt"
)

func main() {
	// fmt.Println(hamming("sadf", "sarg"))
	fmt.Println(hamming("Hello, 世界", "Hello, e界"))
}

func hamming(str1 string, str2 string) int {
	// inti := 1
	// floti := float(ini)
	// var intSice []int = []int{1, 2, 4}
	// intSice := []int{1, 2, 4}
	// var foo rune
	// foo = 'Ü'
	// runesi2 := []rune{'a', 'ä'}
	runesi1 := []rune(str1)
	runesi2 := []rune(str2)
	if len(runesi1) != len(runesi2) {
		return -1
	}
	result := 0
	for position, vRune := range runesi1 {
		// if string(vRune) != string([]rune(str2)[position]) {
		if vRune != runesi2[position] {
			result++
		}
	}

	return result
}
