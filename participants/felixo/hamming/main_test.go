package main

import "testing"

// var fixture = []struct{

// }

// var testCases = []struct {
// 	ham1        string
// 	ham2        string
// 	expected    int
// 	description string
// }{
// 	{"abb", "abb", 0, "equal hamming strings"},
// }

func TestHamming(t *testing.T) {
	var input1 string
	var input2 string
	var result int
	var expected int

	input1 = "foo"
	input2 = "ba3"
	expected = 3
	result = hamming(input1, input2)
	if result != expected {
		t.Errorf("unequal hamming strings, give %s %s, expected %d, result %d", input1, input1, expected, result)
	}

	// result = hamming("foo", "foo")
	// if result != 0 {
	// 	t.Error("equal hamming strings expected")
	// }

	// t.Error("TEST_FAILED")

	// for i, v := range testCases {
	// 	t.Run(func(t *testing.T) {

	// 	})
	// }
}
