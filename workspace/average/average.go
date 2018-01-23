package main

import "fmt"

func average (xs [] float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func main () {
	xs := []float64{98, 87, 88, 79, 91, 86}
	fmt.Println(average(xs))
}