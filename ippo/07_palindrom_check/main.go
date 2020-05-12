package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(isPalindrome("racecar"))
	fmt.Println([]rune{'a', 'b', 'c', 'd'}, invert([]rune{'a', 'b', 'c', 'd'}))
	// fmt.Println(isPalindrome("aa"))
	// fmt.Println([]rune{'a', ' ', 'b'})
	// fmt.Println(getRidIgnoredRunes([]rune{'a', ' ', 'b'}))

}

func isPalindrome(maybePalindrome string) bool {
	return isPalindromeStartegyOne(maybePalindrome)
}

func isPalindromeStartegyOne(maybePalindrome string) bool {
	// getRidIgnoredRunes(maybePalindrome)
	maybePalindromeRunes := getRidIgnoredRunes([]rune(strings.ToLower(maybePalindrome)))
	n := len(maybePalindromeRunes)

	//filter out from maybePalindromeRunes all ignored runes
	// something := somefn(maybePalindromeRunes, ignoredRunes)

	//0 <= i <= n
	for i := 0; i < n; i++ {
		if maybePalindromeRunes[i] == ' ' {
			continue
		}

		// fmt.Printf("n: %d, i: %d, n-1: %d, n-i: %d, n-1-i: %d \n", n, i, n-1, n-i, n-1-i)
		if maybePalindromeRunes[i] != maybePalindromeRunes[(n-1)-i] {
			return false
		}
	}

	return true
}

// example:
// elementsToFilterOut := []rune{' ', ','}
// func filterBySlice(unfiltered []rune, elementsToFilterOut []rune) filtered []rune {
// }

func getRidIgnoredRunes(maybePalindrome []rune) []rune {
	newBucket := []rune{}

	for _, v := range maybePalindrome {
		if v == ' ' {
			continue
		}

		newBucket = append(newBucket, v)
	}

	return newBucket
}

// for initializiation; break condition; continuation statement {}

// true loop!!!!
// for true {
// 	if someCondition {
// 		break
// 	}
// }

// i,  0<= i <= n, 0==n, 1=n-i, 2=n-i, ...
//n := len(maybePalindrim)

// 2nd
// take string, print it out
// take the string, invert it
// print it out again, if equal, isPalindrom

func isPalindromeStartegyTwo(maybePalindrome string) bool {
	maybePalindromeRunes := getRidIgnoredRunes([]rune(strings.ToLower(maybePalindrome)
	// var1 := cleanUpStringAndMaybeTurnItIntoRune(maybePalindrome)
	// va
}

func invert(notInverted []rune) (inverted []rune) {

	l := len(notInverted)
	// for i := 0; i < l; i++ {
	// for i, _ := range notInverted {
	// 	inverted = append(inverted, notInverted[(l-1)-i])
	// }

	for i := l - 1; i >= 0; i-- {
		inverted = append(inverted, notInverted[i])
	}

	return inverted
}

// for i,v := range notInverted{
// 	//how do I get with the index 0 the element of the (last index)
// 	//how do I get with index 1 the elemnt of the (last index) ?????
// 	//how do I get with index 2 the elemnt of the ????? index

// 	// inverted := append(inverted, ... )
// }
