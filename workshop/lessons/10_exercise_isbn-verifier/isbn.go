package main

import (
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// fmt.Println(IsValidISBN("3-598-21508-8"))
	// fmt.Println(IsValidISBN("3-598-21508-X"))
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	// remove all '-'
	isbn = strings.Replace(isbn, "-", "", -1)

	isbnRunes := []rune(isbn)

	//check length == 10
	if len(isbnRunes) != 10 {
		return false
	}

	//check every char (without checking the last) isDigit
	for i, isbnRune := range isbnRunes {
		if i == len(isbnRunes)-1 {
			break
		}

		if !unicode.IsDigit(isbnRune) {
			return false
		}
	}

	// check last char is digit or x
	last := isbnRunes[len(isbnRunes)-1]
	if !(unicode.IsDigit(last) || last == 'X' || last == 'x') {
		return false
	}

	//transform isbn toNumeric representation
	var numericISBN [10]int
	for i, v := range isbnRunes {
		if i == len(isbnRunes)-1 && (v == 'X' || v == 'x') {
			numericISBN[i] = 10
			continue
		}

		numeric, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		numericISBN[i] = numeric
	}

	// check numeric isbn logic is valid
	var betweenResult int = 0
	for i := range numericISBN {
		if i == 9 {
			betweenResult = betweenResult + numericISBN[i]
			continue
		}
		betweenResult = betweenResult + (numericISBN[i] * (10 - i))
	}

	result := betweenResult % 11

	return result == 0
}
