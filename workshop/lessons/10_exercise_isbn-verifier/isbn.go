package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

//v1
//go run isbn.go "3-598-21508-8"
//
//given: "3-598-21508-8"
//is isbn: true
//
//v2
//go run isbn.go -f /path/to/file
//
//input
//3-598-21508-8
//3-598-21508-8
//3-598-21508-8
//
//output:
//isbn; is valid;
//3-598-21508-8; true
//3-598-21508-8; true
//3-598-21508-8; true
//
func main() {
	f := flag.Bool("f", false, "Input argument is a CSV file")
	flag.Parse()

	if flag.NArg() != 1 {
		log.Fatal("Argument count mismatch")
	}

	if !*f {
		isbnInput := flag.Arg(0)
		fmt.Printf("%s: %v\n", isbnInput, IsValidISBN(isbnInput))
		return
	}

	csvfile, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	// r := csv.NewReader(csvfile)
	r := csv.NewReader(bufio.NewReader(csvfile))

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		isbnInput := record[0]
		fmt.Printf("%s: %v\n", isbnInput, IsValidISBN(isbnInput))

	}

	//	if (flag.)
	// isbnInput := flag.Arg(0)

	// if len(argsWithoutExecutable) > 1 {
	// 	fmt.Errorf()
	// }

	// arg := os.Args[3]

	// fmt.Println(argsWithProg)
	// fmt.Printf("given: %s\n", isbnInput)
	// fmt.Printf("is valid ISBN: %v", IsValidISBN(isbnInput))
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
