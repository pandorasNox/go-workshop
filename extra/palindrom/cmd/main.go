package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pandorasNox/go-workshop/extra/palindrom"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	if !palindrom.Is(string(data)) {
		fmt.Printf("invalid")
		os.Exit(1)
	}

	fmt.Printf("valid")
}
