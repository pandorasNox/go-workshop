package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// func main() {
// 	// fmt.Println(IsValidISBN("3-598-21508-8"))
// 	// fmt.Println(IsValidISBN("3-598-21508-X"))

// 	argsWithoutExecutable := os.Args[1:]

// 	if len(argsWithoutExecutable) == 0 {
// 		log.Fatalln("missing argument(s)")
// 	}

// 	if len(argsWithoutExecutable) != 1 {
// 		log.Fatalln("to manny arguments")
// 	}

// 	maybeISBN := argsWithoutExecutable[0]

// 	fmt.Printf("given: \"%s\"\n", maybeISBN)
// 	fmt.Println("is isbn:", IsValidISBN(maybeISBN))
// }

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}

// IsValidISBN returns true for valid isbn's | otherwise - false
func IsValidISBN(isbn string) bool {

	// isbn = dropHyphen(isbn)
	isbn = strings.ReplaceAll(isbn, "-", "")

	isbnNumeric, err := toNumeric(isbn)
	if err != nil {
		return false
	}

	return isValidNumericISBN(isbnNumeric)
}

func toNumeric(isbn string) (result []int, err error) {

	for pos, char := range isbn {
		switch {
		case unicode.IsLetter(char) && (char != 'X' || pos != 9):
			err = errors.New("invalid character")
			return
		case char == 'X':
			result = append(result, 10)
		default:
			i, _ := strconv.Atoi(string(char))
			result = append(result, i)
		}
	}
	return
}

func isValidNumericISBN(isbn []int) bool {
	if len(isbn) != 10 {
		return false
	}

	var pool int
	for idx, value := range isbn {
		pool += int(math.Abs(float64(idx)-10)) * value
	}
	result := pool % 11

	return result == 0
}
