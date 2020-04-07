package main

import (
	"fmt"
)

var a int = 6
var b bool = false

type person struct{
	hasName bool
	age int
}

func main() {
	fmt.Println("")
	fmt.Println("Hello Structs:")
	fmt.Println("")

	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	// var c byte = 107
	// fmt.Println("c:", c)

	// var d = []byte{107, 108, 109}
	// fmt.Println("d:", d)

	//struct
	// type person struct{
	// 	hasName bool
	// 	age int
	// }

	// var st person = person{
	// 	hasName: true,
	// 	age:107,
	// }
	// fmt.Println("st:", st)
	// fmt.Println("st.age:", st.age)
	// //changes
	// st.age = 56
	// fmt.Println("st.age:", st.age)

	//test cases example:
	type testData struct {
		description string
		// in1 string
		// in2 string
		input [2]string
		//everything we expect:
		expectedResult int
		wantErr bool
	}
	// type listOfTests []testData

	// var testCases = listOfTests{
	//var testCases = []testData{
	//	testData{
	//		description: "hello world",
	//		// in1 string
	//		// in2 string
	//		input: [2]string{"AA", "AB"},
	//		// input struct {
	//		// 	in1 string
	//		// 	in2 int
	//		// }
	//		expectedResult: 1,
	//		wantErr: false,
	//	},
	//	testData{
	//		description: "other",
	//		input: [2]string{"AA", "AC"},
	//		expectedResult: 1,
	//		wantErr: false,
	//	},
	//}

	// fmt.Println("testCases[]:", testCases)
	// fmt.Println("testCases[0]:", testCases[0])

	var testCases = []struct {
		in1         string
		in2         string
		expected    int
		wantErr     bool
		description string
	}{
		{"AA", "AA", 1, false, "different hamming inputs, result 1"},
		{"", "", 1, false, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
		{"AAA", "AB", 1, true, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
		{"AA", "AB", 1, false, "different hamming inputs, result 1"},
	}
	fmt.Println("testCases[]:", testCases)
}

// func TestLen(t *testing.T) {
// 	inS1 := "A"
// 	inS2 := ""
// 	observedRes, observedErr := calcHamming(inS1, inS2)
// 	expectedRes := 0
// 	wantErr := true
// 	if observedRes != expectedRes {
// 		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
// 	}
// 	hasObservedErr := (observedErr != nil)
// 	if wantErr != hasObservedErr {
// 		t.Errorf("Can not calculate Hamming distance between '%s' / '%s' because observed value '%d' & wantErr '%t', is not expected val %d & has expected error %t ", inS1, inS2, observedRes, (observedErr != nil), expectedRes, wantErr)
// 	}
// }

// //josn example
// {
// 	"name": "Jason",
// 	"age": 16,
// }
// {
// 	"someSubObject": {
// 		"name"; "foo",
// 		"age"; 56
// 	}
// }
// //yaml
// name: Jason
// age 16
