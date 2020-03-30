package main

import (
	"testing"
)

func TestCheckLen(t *testing.T) {
	inS1 := "A"
	inS2 := ""
	observedRes, observedErr := calcHamming(inS1, inS2)
	expectedRes := 1
	hasExpectedError := false
	if observedRes != expectedRes {
		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & hasObservedError '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, hasExpectedError)
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
