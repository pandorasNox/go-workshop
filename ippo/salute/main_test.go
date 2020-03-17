package main

import (
	"testing"
)

func TestSalute(t *testing.T) {
	//	var inputLastname string = "Smith"
	inputLastname := "Smith"
	inputGender := "male"
	expected := "Mr. Smith"
	givenResult := salute(inputLastname, inputGender)
	if givenResult != expected {
		t.Errorf("Salute failed, given: '%s' / expected: '%s'", givenResult, expected)
	}
}
