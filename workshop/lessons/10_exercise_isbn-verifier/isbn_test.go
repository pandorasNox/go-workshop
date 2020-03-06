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
		{"isbn without any numbers", args{"no isbn"}, false},
		{"too long isbn", args{"35982150881"}, false},
		{"too short isbn", args{"359821508"}, false},
		{"valid isbn with pure numbers", args{"3598215088"}, true},
		{"valid isbn with dashes", args{"3-598-21508-8"}, true},
		{"valid isbn with spaces", args{"3 598 21508 8"}, true},
		{"valid isbn with spaces and dashes", args{"3 598-21508-8"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidISBN(tt.args.isbn); got != tt.want {
				t.Errorf("IsValidISBN() = %v, want %v", got, tt.want)
			}
		})
	}
}
