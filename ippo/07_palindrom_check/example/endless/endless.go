package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start (before loop)")
	count := 0
	for {
		if count == 25 {
			break
		}
		count++

		fmt.Println("racecar")
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("end (after loop)")

}
