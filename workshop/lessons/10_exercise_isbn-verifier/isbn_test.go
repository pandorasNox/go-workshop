package main

import "testing"

var testCases = []struct {
	givenISBN   string
	expected    bool
	description string
}{
	{"", false, "invalid isbn - empty"},
	{"12", false, "invalid isbn - too short"},
	{"123456789", false, "invalid isbn - less than 10 characters"},
	{"123-456-789", false, "invalid isbn - less than 10 characters, with dashes"},
	{"3-598-21507-x", true, "valid isbn - small letter"},
	{"123456789X", true, "invalid isbn"},
	{"0-8436-1072-7", true, "valid isbn"},
	{"3-598-21508-9", false, "invalid isbn - check digit"},
	{"0-306-40615-2", true, "valid isbn - just numbers"},
	{"3-598-21507-A", false, "check digit is a character other than X"},
}

func TestIsValidISBN(t *testing.T) {
	for _, v := range testCases {
		if IsValidISBN(v.givenISBN) != v.expected {
			t.Error(
				"TEST_FAILED, description: '", v.description, "' ",
				"input: '", v.givenISBN, "' ",
				"evaluated to: '", IsValidISBN(v.givenISBN), "' ",
				"expected: '", v.expected, "' ",
			)
		}
	}
}
