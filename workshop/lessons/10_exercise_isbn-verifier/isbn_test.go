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
		{"valid", args{"3-598-21508-8"}, true},
		{"valid with X", args{"3-598-21507-X"}, true},

		{"invalid length", args{""}, false},
		{"invalid length", args{"no isbn"}, false},
		{"invalid length", args{"3-598-21508-10"}, false},
		{"invalid check digit", args{"3-598-21508-A"}, false},
		{"invalid", args{"3-598-21508-X"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidISBN(tt.args.isbn); got != tt.want {
				t.Errorf("IsValidISBN() = %v, want %v", got, tt.want)
			}
		})
	}
}
