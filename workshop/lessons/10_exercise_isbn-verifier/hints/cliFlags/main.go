package main

import (
	"flag"
	"fmt"
)

// somecmd -f -g --hgdfdf sadf asfdasd asdasd

// go run main.go -smt=flase -file

func main() {
	isSmt := flag.Bool("smt", true, "is smt")
	file := flag.String("file", "/foo", "path to file")
	flag.Parse()

	// flag.NArg()
	// flag.Args()

	fmt.Println(*file, *isSmt)
}
