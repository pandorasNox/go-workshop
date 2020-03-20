package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

//go run isbn.go "3-598-21508-8"
//
//given: "3-598-21508-8"
//is isbn: true
func main() {
	// fmt.Println(IsValidISBN("3-598-21508-8"))
	// fmt.Println(IsValidISBN("3-598-21508-X"))

	callArgs := os.Args[1:]

	// if len(callArgs) != 1 {
	// 	log.Fatal("Argument count mismatch")
	// }

	f := flag.Bool("f", false, "Input argument is a CSV file")
	flag.Parse()
	log.Println(*f)

	if *f {
		log.Println("input is a file")
	}

	//	if (flag.)
	isbnInput := callArgs[0]

	// if len(argsWithoutExecutable) > 1 {
	// 	fmt.Errorf()
	// }

	// arg := os.Args[3]

	// fmt.Println(argsWithProg)
	fmt.Printf("given: %s\n", isbnInput)
	fmt.Printf("is valid ISBN: %v", IsValidISBN(isbnInput))
	// fmt.Println(arg)

	// s := "ab"
	// fmt.Println(s[1])
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
