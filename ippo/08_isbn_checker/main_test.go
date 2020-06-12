package main

import (
	"reflect"
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	// func IsValidISBN(isbn string) bool {

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

	for _, testCase := range testCases {
		givenRes := IsValidISBN(testCase.input)
		// observedRes := IsValidISBN(testCase.input)
		if givenRes != testCase.expectedRes {
			t.Error("TEST_FAILED")
		} else {
			t.Log("Test_Pass")
		}
	}

}

//case 1 ISBN 0-8436-1072-7 true
//case 2 ISBN 2-uzg5-521x-i false

//description string //"valid isbn" "invalid isbn" "valid isbn with only didgits" "valid isbn with an X"

func Test_vectorProduct(t *testing.T) {
	type args struct {
		vec1 []int
		vec2 []int
	}
	tests := []struct {
		name       string
		args       args
		wantResVec []int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "multiply two equal vectors",
			args: args{
				vec1: []int{1, 2},
				vec2: []int{2, 1},
			},
			wantResVec: []int{2, 2},
			wantErr:    false,
		},
		{
			"multiply two equal vectors",
			args{[]int{1, 2}, []int{2, 1}},
			[]int{2, 2},
			false,
		},
		{
			name: "multiply two differnet lenght vectors",
			args: args{
				vec1: []int{1, 2},
				vec2: []int{2, 1, 1},
			},
			wantResVec: []int{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResVec, err := vectorProduct(tt.args.vec1, tt.args.vec2)
			if (err != nil) != tt.wantErr {
				t.Errorf("vectorProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResVec, tt.wantResVec) {
				t.Errorf("vectorProduct() = %v, want %v", gotResVec, tt.wantResVec)
			}
		})
	}
}
