package main

import (
	"os/exec"
	"strings"
	"testing"
)

var testCases = []struct {
	input       string
	expected    bool
	description string
}{
	{"", false, "invalid, empty"},
	{"0-8436-1072-7", true, "valid"},
	{"0843610727", true, "valid"},
	{"0-8-4-3-6-1-0-7-2-7", true, "valid"},
	{"3-598-21507-X", true, "valid isbn number with a check digit of 10 upper case"},
	{"3-598-21507-x", true, "valid isbn number with a check digit of 10 lower case"},
	{"3-598-21507-A", false, "invalid digit"},
	{"3-598-21507", false, "invalid number of digits < 10"},
	{"3-598-21507-1-2", false, "invalid number of digits > 10"},
	{"3-598-21507-1", false, "invalid checksum"},
	{"3-5d8-21507-1", false, "invalid checksum"},
	{"3-598-21507-1", false, "invalid checksum"},

	{"3-598-21508-8", true, "valid isbn number"},
	{"3-598-21508-9", false, "invalid isbn check digit"},
	{"3-598-21507-X", true, "valid isbn number with a check digit of 10"},
	{"3-598-21507-A", false, "check digit is a character other than X"},
	{"3-598-P1581-X", false, "invalid character in isbn"},
	{"3-598-2X507-9", false, "X is only valid as a check digit"},
	{"3598215088", true, "valid isbn without separating dashes"},
	{"359821507X", true, "isbn without separating dashes and X as check digit"},
	{"359821507", false, "isbn without check digit and dashes"},
	{"3598215078X", false, "too long isbn and no dashes"},
	{"00", false, "too short isbn"},
	{"3-598-21507", false, "isbn without check digit"},
	{"3-598-21515-X", false, "check digit of X should not be used for 0"},
	{"", false, "empty isbn"},
	{"134456729", false, "input is 9 characters"},
	{"3132P34035", false, "invalid characters are not ignored"},
	{"98245726788", false, "input is too long but contains a valid isbn"},
}

/**
 * validation unit test
 */
func TestIsValidISBN(t *testing.T) {

	//func IsValidISBN(isbn string) bool {
	for _, testCase := range testCases {
		observed := IsValidISBN(testCase.input)
		if observed != testCase.expected {
			t.Errorf("input='%s', observed='%t', expected='%t' - description: %s", testCase.input, observed, testCase.expected, testCase.description)
		}
	}

}

/**
 * cli integration test
 */
func TestCli(t *testing.T) {
	// content validation for single isbn
	for _, testCase := range testCases {
		observed, err := exec.Command("go", "run", "isbn.go", testCase.input).Output()
		if err != nil {
			t.Errorf("CLI exec command failed for input='%s'",testCase.input)
			continue
		}
		expected := "the given input '" + testCase.input + "' to be expected an isbn is: "
		if testCase.expected {
			expected += "valid"
		} else {
			expected += "invalid"
		}
		if strings.TrimSpace(string(observed)) != expected {
			t.Errorf("CLI input='%s', observed='%s\n', expected='%s' - description: %s", testCase.input, observed, expected, testCase.description)
		}
	}
}
