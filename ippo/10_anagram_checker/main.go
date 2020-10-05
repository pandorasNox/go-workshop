package main

import (
	"fmt"
	"strings"
)

func main() {
	r := IsAnagram("aba", "baa")
	fmt.Println(r)
}

// IsAnagram return true if maybeAnagramOfWord is an anagram of word, otherwise returns false
func IsAnagram(word, maybeAnagramOfWord string) bool {
	// replace uppercase letters with low
	// remove white spaces from word and maybeAnagramOfWord
	// handle the special characters (runes...)
	lowCleanWord := strings.ToLower(word)
	cleanedWord := strings.ReplaceAll(lowCleanWord, " ", "")
	cleanedWord = strings.ReplaceAll(cleanedWord, "-", "")
	cleanedWord = strings.ReplaceAll(cleanedWord, "'", "")
	lowCleanedmaybeAnagramOfWord := strings.ToLower(maybeAnagramOfWord)
	cleanedmaybeAnagramOfWord := strings.ReplaceAll(lowCleanedmaybeAnagramOfWord, " ", "")
	cleanedmaybeAnagramOfWord = strings.ReplaceAll(cleanedmaybeAnagramOfWord, "-", "")
	cleanedmaybeAnagramOfWord = strings.ReplaceAll(cleanedmaybeAnagramOfWord, "'", "")

	// check if the len of word and maybeAnagramOfWord are equal
	if len(cleanedWord) != len(cleanedmaybeAnagramOfWord) {
		return false
	}

	wordLetterMap := map[rune]uint8{}

	for _, v := range cleanedWord {
		// if _, keyExist := wordLetterMap[v]; keyExist {
		// 	// wordLetterMap[v] = wordLetterMap[v] + 1
		// 	wordLetterMap[v]++
		// } else {
		// 	// wordLetterMap[v] = 1
		// 	// wordLetterMap[v] = wordLetterMap[v] + 1
		// 	wordLetterMap[v]++
		// }

		wordLetterMap[v]++
	}

	maybeLetterMap := map[rune]uint8{}

	for _, v := range cleanedmaybeAnagramOfWord {
		maybeLetterMap[v]++
	}

	if len(wordLetterMap) != len(maybeLetterMap) {
		return false
	}

	for key, value := range wordLetterMap {
		if _, keyExist := maybeLetterMap[key]; !keyExist {
			return false
		}

		if value != maybeLetterMap[key] {
			return false
		}
	}

	return true
}

//abc & abc
//a -> a
//=> bc & bc
//...

//aba & abc

//abc & aba

//using a map
// key => val
//map[rune]true
//a => true
//b => false
//c => false

//using a map
// key => val
//map[rune]uint
//example abc
//a => 1
//b => 1
//c => 1
//example aba
//a => 2
//b => 1

//clever solution
// iterate over maybe and subtract 1 by found
// check if 0 is everywhere in the map

//rune[] = in32[]
//we can sort both
//then compare both, if there is one different we can return early with false

//abc & abc
// takeout a from slice of runeWord and search for in maybeAWord, if found take it out there too
//// otherwise return false

// remove white spaces from word and maybeAnagramOfWord
// handle the special characters (runes...)
// check if the len of word and maybeAnagramOfWord are equal
// check if the elements of the "word" exist at "maybeAnagramOfWord"
//// if letter not found, return false
//// if letter found take it out somehow
//////???

//[A, B]
//[10,0]
//[C, C]
//[5,5]

// [89,112,50]
// [50,89,112]
