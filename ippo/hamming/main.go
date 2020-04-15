package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(calcHamming("AA", "A"))
}

func calcHamming(inS1 string, inS2 string) (int, error) {
	// // len()
	// if len(inS1) == len(inS2) {
	// 	return 0, nil
	// } else if len(inS1) != len(inS2) {
	// 	return -1, errors.New("FAILED")
	// }

	//handle runes
	// inS1Ru := []rune("Üm")
	// fmt.Println(len(inS1Ru), inS1Ru[0], inS1Ru[1], string(inS1Ru[0]), string(inS1Ru[1]))
	inS1Ru := []rune(inS1)
	inS2Ru := []rune(inS2)

	if len(inS1Ru) != len(inS2Ru) {
		return -1, errors.New("FAILED")
	}

	//calc
	difference := 0
	for i := range inS1Ru {
		if inS1Ru[i] != inS2Ru[i] {
			difference++
		}
	}

	return difference, nil
}

// else {
// 	difference = difference + 0
// 	difference += 0
// }
