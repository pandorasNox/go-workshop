package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

// flag.NArg()
// flag.Args()
func main() {
	flag.Parse()

	//argsWithProg := os.Args

	if flag.NArg() > 1 {
		log.Fatal("Argument count mismatch, only one argument allowed")
	}

	input := flag.Args()[0]

	// // fmt.Println("The input '", input, "' is valid ISBN", IsValidISBN(input))
	fmt.Printf("The input '%s' is valid ISBN: %t \n", input, IsValidISBN(input))
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	isbn1 := strings.Replace(isbn, "-", "", -1)
	// fmt.Println(isbn1)
	isbn2 := []rune(isbn1)

	if len(isbn2) != 10 {
		fmt.Println("to long")
		return false
	}

	//check first 9 symbols are digits
	for i := 0; i < 9; i++ {
		if !unicode.IsDigit(isbn2[i]) {
			fmt.Println("the first 9 has to be digits")
			return false
		}
	}

	//check last symbol is digit or 'x' or 'X'
	allowedSymbols := []rune{'x', 'X'}
	lastElem := len(isbn2) - 1
	if unicode.IsDigit(isbn2[lastElem]) == false && find(isbn2[lastElem], allowedSymbols) == false { // == false can be replaced !find.
		//if !(unicode.IsDigit(isbn2[lastElem]) || find(isbn2[lastElem], bla)) {

		//fmt.Println("last elem:", isbn2[lastElem], unicode.IsDigit(isbn2[lastElem]))
		fmt.Println("the last has to be digits or x or X")
		return false
	}

	//transform to int
	isbnNumeric := []int{}
	for i, _ := range isbn2 {

		if i == 9 {
			if isbn2[i] == 'x' || isbn2[i] == 'X' {
				//isbn2[i] == 10
				//isbnNumeric[i] = 10
				isbnNumeric = append(isbnNumeric, 10)
				break
			}
		}

		//transform the singleLetter into a numeric (int) value
		//'3' => 3
		singleNumeric, err := strconv.Atoi(string(isbn2[i]))
		if err != nil {
			return false
		}

		isbnNumeric = append(isbnNumeric, singleNumeric)
		//isbnNumeric = append(isbnNumeric, ??????)
	}

	// fmt.Println("numeric: ", isbnNumeric)

	weightsVector := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	finalVector, err := vectorProduct(isbnNumeric, weightsVector)
	if err != nil {
		return false
	}

	// fmt.Println("finalVector: ", finalVector)

	var sumFinalVector int
	for i := 0; i < len(finalVector); i++ {
		sumFinalVector = sumFinalVector + finalVector[i]
	}

	// fmt.Println("sumFinalVector", sumFinalVector)

	// if sumFinalVector % 11 != 0 {
	// 	return false
	// }

	// return true

	return sumFinalVector%11 == 0
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

func vectorProduct(vec1, vec2 []int) (resVec []int, err error) {
	if len(vec1) != len(vec2) {
		return []int{}, errors.New("vectors have not the same length")
	}

	resVec = make([]int, len(vec1), len(vec1))

	// vec1 [x1, ,x2, x3, ...] x vec2 [y1, y2, y3, ...] == resVec [x1*y1, x2*y2, x3*y3, ...]
	for i := 0; i < len(vec1); i++ {
		resVec[i] = vec1[i] * vec2[i]
		//isbnNumeric = append(isbnNumeric, singleNumeric)

	}

	return resVec, nil
}

// allow "-" in the string, but ignore for calculation   (√)
//// how to remove a char from a string
//// how to filter out.....
// check length is 10									 (√)
// accepted only 0-9 and X,x at last position            (√)
// transform the single characters of the string to slice of int (vector1: one dimension matrix)  (√)
//// last position is X,x replaced by 10				 (√)
// multiply vector1 with vector2 (10,9...,2,1) --> vector3
// sum all elements of vector3
// when (sum % 11 == 0) is valid
// when (sum % 11 != 0) is invalid
