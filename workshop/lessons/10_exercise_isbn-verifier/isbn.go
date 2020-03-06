package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	if isbn == "" {
		return false
	}

	if len(isbn) > 10 {
		return false
	}

	return true
}
