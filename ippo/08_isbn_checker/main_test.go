package main

import "testing"

func TestIsValidISBN(t *testing.T) {
	testCases := []struct {
		input       string
		expectedRes bool
		description string
	}{
		{"0-8436-1072-7", true, "valid isbn"},
		{"2-uzg5-521x-i", false, "invalid isbn because has forbitten characters on forbitten position"},
		{"3-598-21507-X", true, "valid isbn number with a check digit of 10 (uppercase X)"},
		{"3-598-21507-x", true, "valid isbn number with a check digit of 10 (lowercase x)"},
	}

	t.Error("TEST_FAILED")
}

//case 1 ISBN 0-8436-1072-7 true
//case 2 ISBN 2-uzg5-521x-i false

//description string //"valid isbn" "invalid isbn" "valid isbn with only didgits" "valid isbn with an X"
