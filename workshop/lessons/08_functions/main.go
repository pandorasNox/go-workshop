package main

import (
	"fmt"
	"math"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}

func average(xs []float64) (float64, error) {
	l := len(xs)
	if l == 0 {
		return 0.0, fmt.Errorf("empty list")
	}

	avg := 0.0
	for _, v := range xs {
		avg += v / float64(l)
	}

	if math.IsInf(avg, 0) {
		return 0.0, fmt.Errorf("too huge")
	}

	return avg, nil
}
