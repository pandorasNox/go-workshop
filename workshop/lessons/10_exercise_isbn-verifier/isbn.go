package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(IsValidISBN("3-598-21508-8"))
}

// IsValidISBN returns ...
func IsValidISBN(isbn string) bool {
	if len(isbn) > 13 || len(isbn) == 0 {
		return false
	}
	a := strings.Split(isbn, "")
	total := 0.0
	weight := 10
	for _, v := range a {
		if v == "-" {
			continue
		}
		if v == "X" || v == "x" {
			v = "10"
		}
		iv, err := strconv.Atoi(v)
		if err == nil {
			total = total + float64(weight*iv)
			weight = weight - 1
		} else {
			return false
		}
	}
	if math.Mod(total, 11) == 0 {
		return true
	}
	return false
}
