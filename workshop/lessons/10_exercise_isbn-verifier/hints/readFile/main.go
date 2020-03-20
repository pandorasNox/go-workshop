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
	// for {
	// 	// Read each record from csv
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf(record[0], record[1])
	// }

	err = nil
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		// Read each record from csv
		fmt.Println(record)
		// fmt.Printf(record[0], record[1])
	}
	// if err == io.EOF {
	// 	break
	// }
	if err == io.EOF {
		log.Fatal(err)
	}

	fmt.Println("finish")
}
