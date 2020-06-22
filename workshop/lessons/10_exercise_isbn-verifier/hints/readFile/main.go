package main

import (
	//"bufio"

	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open the file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	// r := csv.NewReader(csvfile)
	r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
		//out: [1, hello, 2020-05]

		fmt.Println(record[0], record[1])
		fmt.Println("")
		//out: 1 hello
	}

	fmt.Println("finish")
}
