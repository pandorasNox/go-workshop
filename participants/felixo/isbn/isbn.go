package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	expectFile := flag.Bool("f", false, "use csv file as input instead")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("argument missmatch, only one argument supported")
	}

	if *expectFile {
		handleCsvMode()
		return
	}

	handleInlineMode()
}

func handleCsvMode() {
	fileName := flag.Arg(0)
	csvfile, err := os.Open(fileName)
	defer csvfile.Close()
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	result := ""
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// validate input
		isValid := IsValidISBN(record[0])
		result += fmt.Sprintf("%s;%t;\n", record[0], isValid)
	}
	fmt.Print(result)
}

func handleInlineMode(){
	isbnCandidate := flag.Arg(0)

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
