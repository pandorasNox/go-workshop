package main

import (
	"flag"
	"fmt"
)

func main() {
	cliFlags := handleFlags()

	fmt.Println(cliFlags)
}

type cliArgs struct {
	flags struct {
		isFile bool
		// anotherFlag string
	}
	args []string
}

func handleFlags() cliArgs {
	isFile := flag.Bool("f", false, "accept path to csv file as argument")
	flag.Parse()

	return cliArgs{
		flags: struct{ isFile bool }{
			isFile: *isFile,
		},
		args: flag.Args(),
	}
}
