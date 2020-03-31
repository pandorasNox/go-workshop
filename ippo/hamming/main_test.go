package main

import (
	"testing"
)

func TestCheckLen(t *testing.T) {
	inS1 := "A"
	inS2 := ""
	observedRes, observedErr := calcHamming(inS1, inS2)
	expectedRes := -1
	wantErr := true
	if observedRes != expectedRes {
		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
	}
	hasObservedErr := (observedErr != nil)
	if wantErr != hasObservedErr {
		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
	}
}

// if givenResult != expected {
// 	t.Errorf("Salute failed, given: '%s' / expected: '%s'", givenResult, expected)
// }

// 5, nil
// -1, error
// -1, error
// -1, error

// 0, nil
// 2, nil
