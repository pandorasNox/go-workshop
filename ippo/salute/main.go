package main

import "fmt"

func main() {
	fmt.Println(salute("Smith", "male"))
	fmt.Println(salute("Jules", "female"))
	fmt.Println(salute("Tzanetakis", ""))

}

func salute(lastname, gender string) string {
	if gender == "male" {
		return "Mr. " + lastname
	} else if gender == "female" {
		return "Ms. " + lastname
	}
	return "stranger"

}
