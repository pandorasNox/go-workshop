package main

import "testing"

func TestIsValidISBN(t *testing.T) {
	type args struct {
		isbn string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"no isbn given", args{""}, false},
		//{"isbn without any numbers", args{"no isbn"}, false},
		{"too long isbn", args{"35982150881"}, false},
		//	{"too short isbn", args{"359821508"}, false},
		//	{"valid isbn with pure numbers", args{"3598215088"}, true},
		//	{"valid isbn with dashes", args{"3-598-21508-8"}, true},
		//	{"valid isbn with spaces", args{"3 598 21508 8"}, true},
		//	{"valid isbn with spaces and dashes", args{"3 598-21508-8"}, true},

		//patricia
		// {"valid", args{"3-598-21508-8"}, true},
		// {"valid with X", args{"3-598-21507-X"}, true},
		// {"valid with less -", args{"3-59821507-X"}, true},
		// {"valid with no -", args{"123456789X"}, true},
		// {"invalid length", args{""}, false},
		// {"invalid length", args{"no isbn"}, false},
		// {"invalid length", args{"3-598-21508-10"}, false},
		// {"invalid check digit", args{"3-598-21508-A"}, false},
		// {"invalid", args{"3-598-21508-X"}, false},

		//tino
		// {"", false, "invalid isbn - empty"},
		// {"12", false, "invalid isbn - too short"},
		// {"123456789", false, "invalid isbn - less than 10 characters"},
		// {"123-456-789", false, "invalid isbn - less than 10 characters, with dashes"},
		// {"3-598-21507-x", false, "invalid isbn - small letter"},
		// {"123456789X", true, "invalid isbn"},
		// {"0-8436-1072-7", true, "valid isbn"},
		// {"3-598-21508-9", false, "invalid isbn - check digit"},
		// {"0-306-40615-2", true, "valid isbn - just numbers"},
		// {"3-598-21507-A", false, "check digit is a character other than X"},

		//thomas

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidISBN(tt.args.isbn); got != tt.want {
				t.Errorf("IsValidISBN() = %v, want %v", got, tt.want)
			}
		})
	}
}

//thomas
// func TestIsValidISBNShouldPassForValidISBN(t *testing.T) {
// 	isbn := "3-598-21508-8"
// 	if !IsValidISBN(isbn) {
// 		t.Errorf("Expected ISBN %s to be valid", isbn)
// 	}
// }

// func TestIsValidISBNShouldPassForValidISBNWithNonIntegerCheckDigit(t *testing.T) {
// 	isbn := "3-598-21507-X"
// 	if !IsValidISBN(isbn) {
// 		t.Errorf("Expected ISBN %s to be valid", isbn)
// 	}
// }

// func TestIsValidISBNShouldFailForTooShortISBN(t *testing.T) {
// 	isbn := "3-59-21508-8"
// 	if IsValidISBN(isbn) {
// 		t.Errorf("Expected ISBN %s to be invalid (too short)", isbn)
// 	}
// }

// func TestIsValidISBNShouldFailForTooLongISBN(t *testing.T) {
// 	isbn := "39-598-21508-8"
// 	if IsValidISBN(isbn) {
// 		t.Errorf("Expected ISBN %s to be invalid (too long)", isbn)
// 	}
// }

// func TestIsValidISBNShouldFailForInvalidCheckDigit(t *testing.T) {
// 	isbn := "3-598-21508-9"
// 	if IsValidISBN(isbn) {
// 		t.Errorf("Expected ISBN %s to be invalid due to wrong check digit", isbn)
// 	}
// }
