package main

import "testing"

// Silent
// Listen

func TestIsAnagram(t *testing.T) {
	type args struct {
		word               string
		maybeAnagramOfWord string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "zero test-case", args: args{"1", ""}, want: false},
		{name: "...", args: args{"a", "ab"}, want: false},
		{name: "...", args: args{"aa", "ab"}, want: false},
		{name: "...", args: args{"ac", "ab"}, want: false},
		{name: "...", args: args{"aa", "aa"}, want: true},
		{name: "first test-case", args: args{"rail safety", "fairy tales"}, want: true},
		{name: "second test-case", args: args{"debit card", "bad credit"}, want: true},
		{name: "third test-case", args: args{"New York Times", "monkeys write"}, want: true},
		{name: "forth test-case", args: args{"Church of Scientology", "ich-chosen goofy cult"}, want: false},
		{name: "fifth test-case", args: args{"McDonald's restaurants", "Uncle Sam's standard rot"}, want: true},
		{name: "sixth test-case", args: args{"coronavirus", "carnivorous"}, want: true},
		{name: "seventh test-case", args: args{"coronavirus", "carnivoroux"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.args.word, tt.args.maybeAnagramOfWord); got != tt.want {
				t.Errorf("IsAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
