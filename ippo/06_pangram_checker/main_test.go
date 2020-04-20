package main

import (
	"testing"
)

func TestPangram(t *testing.T) {
	testCases := []struct {
		maybePangram string
		expectedRes  bool
	}{
		{"the quick brown fox jumps over the lazy dog", true},
		{"Shaw, those twelve beige hooks are joined if I patch a young, gooey mouth", true},
		{"Are those shy Eurasian footwear, cowboy chaps, or jolly earthmoving headgear", true},
		{"My ex pub quiz crowd gave joyful thanks", true},
		{"A wizard’s job is to vex chumps quickly in fog", true},
		{"Hello world", false},
		{"Happy day all day", false},
		{"the quick brown foe jumps over the lazy dog", false},
		{"Mexican History is great", false},
		{"I never go to supermarket on Saturday", false},
		{"I never go to supermarket on Saturday", false},
		{"My ex pub quiss crowd gave joyful thanks", false},
	}

	//func calcHamming(inS1 string, inS2 string) (int, error) {
	//func isPangram(maybePangram string) bool {}

	for _, singleCase := range testCases {
		//isPangram(singleCase.maybePangram)
		observedRes := isPangram(singleCase.maybePangram)
		if observedRes != singleCase.expectedRes {
			t.Errorf("Failed isPangram: input:'%s', expected:'%t', observed:'%t' does not match", singleCase.maybePangram, singleCase.expectedRes, observedRes)
			//test failed, reason: given %input and %observedRes the %expectedRes doesn't match}
		} else {
			t.Log("succeeded")
		}
	}
}

//case 0
// input: "the quick brown fox jumps over the lazy dog"
// expected result: true (it's a pangram)

//case 1
// input: "Shaw, those twelve beige hooks are joined if I patch a young, gooey mouth"
// expected result: true (it's a pangram)

//case 2
//"Are those shy Eurasian footwear, cowboy chaps, or jolly earthmoving headgear"
// expected result: true (it's a pangram)

//case 3
//input: "My ex pub quiz crowd gave joyful thanks"
//expected result: true

//case 4
//input: "A wizard’s job is to vex chumps quickly in fog"
//expected result: true

//case 5
//input: "Hello world"
//expected result: false

//case 6
//input: "Happy day all day"
//expected result: false

//case 7
// in: "the quick brown foe jumps over the lazy dog"
// res: false

//case 8
//input: "Mexican History is great"
//expected result: false

//case 9
//input: "I never go to supermarket on Saturday"
//expected result: false

//case 10
//input: "My ex pub quiss crowd gave joyful thanks"
//expected result: false
