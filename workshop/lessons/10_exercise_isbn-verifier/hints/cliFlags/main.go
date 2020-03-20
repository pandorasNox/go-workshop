package main

import (
	"flag"
	"fmt"
)

func main() {
	isSmt := flag.Bool("smt", true, "is smt")
	file := flag.String("file", "/foo", "path to file")
	flag.Parse()

	fmt.Println(*file, *isSmt)
}
