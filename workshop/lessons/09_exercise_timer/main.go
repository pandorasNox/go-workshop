package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	duration, err := time.ParseDuration("10s")

	if err != nil {
		panic(err)
		os.Exit(2)
	}

	// startTime := time.Now()
	// endTime := startTime.Add(duration)

	i := int64(1)
	ticker1 := time.NewTicker(1 * time.Second)
	for _ = range ticker1.C {
		i++
		if i > int64(duration/time.Second) {
			ticker1.Stop()
			fmt.Println("Done")
			break
		}
	}
}
