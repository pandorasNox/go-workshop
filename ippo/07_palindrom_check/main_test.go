package main

import (
	"testing"
)

var allowedAlphabet = []rune{'a', '1'}
var ignoreAlphabet = []rune{' ', ',', '`'}

func TestPalindrom(t *testing.T) {
	testCases := []struct {
		maybePalindrome string
		expectedRes     bool
	}{
		// {"A", true}, //??????
		{"AA", true},
		{"AB", false},
		{"aa", true},
		{"ABA", true},
		{"aba", true},
		{"AmoreRoma", true},
		{"Amore Roma", true},
		{"A nut for a jar of tuna", true},
		{"Was it a car or a cat I saw", true},
		{"Murder for a jar of red rum", true},
		{"0101010101010", true},
		{"0808090908080", true},
		{"A99A", true},
		{"A919A", true},
		{"lagerregal", true},
		{"madam", true},
		{"Murder for a jar of rotten rum", false},
		{"As I walk through the valley of the shadow of death", false},
		{"A happy day on earth", false},
		{"I have a dream that my dream will no come true", false},
		{"08148194614562", false},
		{"6982371300821", false},
	}

	for _, singleCase := range testCases {
		obsevedRes := isPalindrome(singleCase.maybePalindrome)
		if obsevedRes != singleCase.expectedRes {
			t.Errorf("failed")
		} else {
			t.Log("succeded")
		}
	}
}

//test cases
// 1. "Amore, Roma" true
// 2. "A nut for a jar of tuna" true
// 3. "Was it a car or a cat I saw" true
// 4. "Murder for a jar of red rum" true
// 5. "0101010101010" true
// 6. "0808090908080" true
// 7. "Murder for a jar of rotten rum" false
// 8. "As I walk through the valley of the shadow of death" false
// 9. "A happy day on earth" false
// 10. "I have a dream that my dream will no come true" false
// 11. "08148194614562" false
// 12. "6982371300821"  false
