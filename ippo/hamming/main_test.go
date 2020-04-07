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
		{"AA", "AB", 0, true},
		{"A", "AA", 1, false},
	}
	observedRes, observedErr := calcHamming(inS1, inS2)
	if observedRes != expectedRes {
		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
	}
	hasObservedErr := (observedErr != nil)
	if wantErr != hasObservedErr {
		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
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
