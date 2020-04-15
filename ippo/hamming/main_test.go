package main

import (
	"testing"
)

func TestLen(t *testing.T) {
	var testCases = []struct {
		inS1        string
		inS2        string
		expectedRes int
		wantErr     bool
	}{
		//test 0:
		{"A", "", -1, true}, // index 0
		//test 1:
		{"AA", "AB", 1, false}, // index 1
		//test 2:
		{"A", "AA", -1, true}, // index 2
		//test 3:
		{"A", "B", 1, false}, // index 3
		//test 4:
		{"ÜE", "UE", 1, false}, // index 4
		//test 5:
		{"", "", 0, false}, // index 5
		//test 6:
		{"FGHFHFZT", "FGHXHUZT", 2, false}, // index 6
		//test 7:
		{"tngsovA", "tngsovX", 1, false}, // index 7
	}

	//index 0
	//inS1 := testCases[0].inS1
	//inS2 := testCases[0].inS2

	for i, singleCase := range testCases {

		observedRes, observedErr := calcHamming(singleCase.inS1, singleCase.inS2)

		if observedRes != testCases[i].expectedRes {
			t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", testCases[i].inS1, testCases[i].inS2, observedRes, (observedErr != nil), testCases[i].expectedRes, testCases[i].wantErr)
		}
		//else if observedRes == testCases[i].expectedRes {
		//	t.Errorf("The Hammimg distance between '%s' / '%s' is %d", testCases[i].inS1, testCases[i].inS2, testCases[i].expectedRes)
		//	t.Log("The Hammimg distance between '%s' / '%s' is %d", testCases[i].inS1, testCases[i].inS2, testCases[i].expectedRes)
		//}
		hasObservedErr := (observedErr != nil)
		if testCases[i].wantErr != hasObservedErr {
			t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", testCases[i].inS1, testCases[i].inS2, observedRes, (observedErr != nil), testCases[i].expectedRes, testCases[i].wantErr)
		}
	}

}

//func TestLen(t *testing.T) {
//	inS1 := "A"
//	inS2 := ""
//	observedRes, observedErr := calcHamming(inS1, inS2)
//	expectedRes := 0
//	wantErr := true
//	if observedRes != expectedRes {
//		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
//	}
//	hasObservedErr := (observedErr != nil)
//	if wantErr != hasObservedErr {
//		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
//	}
//}

//matrix := [][]int{[]int{}, []int{},}
//matrix := make([][]int, 2)

//TestLen2
// inS1 := "A"
// inS2 := "AA"

//Test???
// inS1 := "A"
// inS2 := "B"

// inS1 := "ÜE"
// inS2 := "UE"

// inS1 := ""
// inS2 := ""

// inS1 := "A"
// inS2 := "A"

// inS1 := "FGHFHFZT"
// inS2 := "FGHXHUZT"

// inS1 := "tngsovA"
// inS2 := "tngsovA"

// inS1 := "tngsovA"
// inS2 := "tnGsovX"

//for i, _ := range testCases {
//execute code
//singleCase := testCases[i]
//testCases[i] === value (strucht)
//testCases[i] === singleCase

//when
//i === 0
//singleCase === {"A", "", 0, true},
//when
//i === 1
//singleCase === {"AA", "AB", 0, true},

//observedRes, observedErr := calcHamming(testCases[i].inS1, testCases[i].inS2)
