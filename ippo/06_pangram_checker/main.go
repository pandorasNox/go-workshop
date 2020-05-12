package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isPangram("the quick brown fox jumps over the lazy dog"))

	//test find:
	//fmt.Println(find('3', []rune{'1', '5'}))
}

func isPangram(maybePangram string) bool {
	// return isPangram1(maybePangram)
	return isPangramStrategyTwo(maybePangram)
}

func isPangramStrategyOne(maybePangram string) bool {
	validationAlphabet := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	// <- create temp memory...
	alphabetMap := make(map[rune]bool)
	for _, v := range validationAlphabet {
		alphabetMap[v] = false
	}
	//a -> false
	//b -> false
	//c -> false
	//...

	for _, singleLetter := range maybePangram {
		singeletterLower := unicode.ToLower(singleLetter)
		if find(singeletterLower, validationAlphabet) {
			//do smth
			//how do we remember that we found that single latter now???
			alphabetMap[singeletterLower] = true
		}
	}

	for _, v := range alphabetMap {
		if v == false {
			return false
		}
	}

	//maybePangram[0] exists on validateAlphabet rune
	//does given rune exist in sliceOfRunes
	//func findLetter(letter rune, alphebet []rune) bool

	//2nd way
	//check every validationAlphabet letter if exists in maybePangram

	// s := []string{"a", "b", "cd"}

	//iterate over maybePangram
	//check if all letters are there
	//

	return true
}

//returns true if the haystack contains the needle
func find(needle rune, haystack []rune) bool {

	//for i, v := range haystack {
	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}

	return false
}

func isPangramStrategyTwo(maybePangram string) bool {
	validationAlphabet := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
		'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}

	// for _, singleLetter := range maybePangram {
	// 	singeletterLower := unicode.ToLower(singleLetter)

	//2nd way
	//check every validationAlphabet letter if exists in maybePangram
	//left paper validationAlphabet
	//right paper is sentance to check is pangramm
	//first letter
	//take a from the validationAlphabet
	//check if the letter exists in the sentance
	//second letter
	//take b from the validationAlphabet
	//check if the letter exists in the sentance
	//...
	// `i` is the current index and `v` is the value of the current index
	for _, validationLetter := range validationAlphabet {
		//validationLetter === a

		//do I found validationLetter in maybePangram

		foundValLetter := false
		for _, maybeMatchingLetter := range maybePangram {
			//maybeMatchingLetter === t

			//a != t ???
			// true -> so
			// return false
			// if validationLetter != unicode.ToLower(maybeMatchingLetter) {
			// 	return false
			// }

			if validationLetter == unicode.ToLower(maybeMatchingLetter) {
				foundValLetter = true
				break
			}
		}

		if foundValLetter == false {
			return false
		}
	}

	return true
}
