package main

import "fmt"

func main() {
	fmt.Println(isPangram("the quick brown fox jumps over the lazy dog"))

	//test find:
	//fmt.Println(find('3', []rune{'1', '5'}))
}

func isPangram(maybePangram string) bool {
	validationAlphabet := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	for i, v := range validationAlphabet {
		if validationAlphabet[i] == maybePangram {
			return true
		}
	}
	//_ = validationAlphabet

	//maybePangram[0] exists on validateAlphabet rune
	//does given rune exist in sliceOfRunes
	//func findLetter(letter rune, alphebet []rune) bool

	//2nd way
	//check every validationAlphabet letter if exists in maybePangram

	// s := []string{"a", "b", "cd"}

	//iterate over maybePangram
	//check if all letters are there
	//

	return false
}

func find(needle rune, haystack []rune) bool {

	//for i, v := range haystack {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}

	return false
}
