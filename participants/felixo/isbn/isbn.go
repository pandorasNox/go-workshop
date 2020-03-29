package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// go run isbn.go "123456789"
// output:
// the given input '123456789' to expected an isbn is: valid

// go run isbn.go -f "path/to/file"
//
// input (isbns.csv):
// 123456789
// 123456789X
// ""
// 3598215088
//
// output:
// 123456789; false;
// 123456789X; true;
// ; false;
// 3598215088; true;
func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		log.Fatal("argument missmatch, only one argument supported")
	}

	isbnCandidate := argsWithoutProg[0]

	isValid := IsValidISBN(isbnCandidate)
	if isValid {
		fmt.Printf("the given input '%s' to be expected an isbn is: valid\n", isbnCandidate)
	} else {
		fmt.Printf("the given input '%s' to be expected an isbn is: invalid\n", isbnCandidate)
	}
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	allowedAlphabetStart := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	//allowedAlphabetLast := []rune{ 'x', 'X'}
	allowedAlphabetLast := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'x', 'X'}
	cleanedInput := []rune{}

	counterRunes := 0
	for _, v := range isbn {
		if v == '-' {
			continue
		}
		if counterRunes < 9 {
			// check allowedAlphabetStart
			if containsRune(v, allowedAlphabetStart) {
				cleanedInput = append(cleanedInput, v)
				counterRunes++
				continue
			}
			return false
		} else if counterRunes == 9 {
			// check allowedAlphabetLast
			if containsRune(v, allowedAlphabetLast) {
				cleanedInput = append(cleanedInput, v)
				counterRunes++
				continue
			}
			return false
		} else {
			return false
		}
	}
	if counterRunes != 10 {
		return false
	}
	// check checksum modulo 11
	return hasValidChecksum(cleanedInput)
}

func containsRune(singleRune rune, allowed []rune) bool {
	for _, v := range allowed {
		if singleRune == v {
			return true
		}
	}
	return false
}

func hasValidChecksum(cleanedInput []rune) bool {
	inBetweenResult := 0
	for i, v := range cleanedInput {
		if v == 'x' || v == 'X' {
			inBetweenResult += 10 * (10 - i)
			continue
		}
		num, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
		inBetweenResult += num * (10 - i)
	}
	return (inBetweenResult % 11) == 0
}
