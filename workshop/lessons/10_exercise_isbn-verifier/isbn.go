package main

import (
	"math"
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

	//check length == 10
	if len(isbn) != 10 {
		return false
	}

	//check every char (without checking the last) isDigit
	for i := 0; i < 9; i++ {
		charByte := isbn[i]
		if !unicode.IsDigit(rune(charByte)) {
			return false
		}
	}

	// check last char is digit or x
	last := isbn[len(isbn)-1]
	if !(unicode.IsDigit(rune(last)) || last == 'X') {
		return false
	}

	//transform isbn toNumeric representation
	var numericISBN [10]int
	for i, v := range isbn {
		if i == 9 && v == 'X' {
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
		betweenResult = betweenResult + (numericISBN[i] * int(math.Abs(float64(i)-10)))
	}

	result := betweenResult % 11

	return result == 0
}
