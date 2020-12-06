package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	normalizedISBN := normalizeISBN(isbn)

	if !isValidLength(normalizedISBN) {
		return false
	}

	if !hasValidChecksum(normalizedISBN) {
		return false
	}

	return true
}

func normalizeISBN(isbn string) string {
	expression := regexp.MustCompile("([^0-9xX])");
	normalizedISBN := expression.ReplaceAllString(isbn, "")

	return normalizedISBN
}

func isValidLength(isbn string) bool {
	return (len(isbn) == 10)
}

func hasValidChecksum(isbn string) bool {
	checksum := calculateChecksum(isbn)

	return (checksum % 11) == 0
}

func calculateChecksum(isbn string) uint16 {
	var sum uint16 = 0
	for position, character := range isbn {
		multiplier := 10 - position
		sum += castDigitToIntAtPosition(character, uint16(position)) * uint16(multiplier)
	}

	return sum
}

func castDigitToIntAtPosition(digit rune, position uint16) uint16 {
	stringDigit := string(digit)
	if (position == 9 && strings.ToLower(stringDigit) == "x") {
		return 10
	}

	intDigit, error := strconv.ParseUint(stringDigit, 10, 16)

	if (error == nil) {
		return uint16(intDigit)
	}

	return 0
}