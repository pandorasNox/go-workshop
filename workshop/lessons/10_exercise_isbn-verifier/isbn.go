package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	// fmt.Println(IsValidISBN("3-598-21508-8"))
	fmt.Println(IsValidISBN("3-598-21508-X"))

}

// IsValidISBN returns ...
func IsValidISBN(maybeIsbn string) bool {

	// isbn = strings.Replace(isbn, "-", "", -1)

	expression := regexp.MustCompile("([^0-9xX])")
	isbn := expression.ReplaceAllString(maybeIsbn, "")

	if len(isbn) != 10 {
		return false
	}

	var isbnNumeric [10]int
	for i, v := range isbn {
		if v == 'x' || v == 'X' {
			isbnNumeric[i] = 10
			continue
		}

		numeric, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}

		isbnNumeric[i] = numeric
	}

	var inBetweenResult int
	for i, num := range isbnNumeric {
		inBetweenResult += (10 - i) * num
	}

	return inBetweenResult%11 == 0
}
