package main

import "fmt"

func getGreater (xs ...int) (greatest int) {
	greatest = -1
	for _, x := range xs {
		if x >= greatest {greatest = x}
	}

	return
}

func main () {
	fmt.Println(getGreater(2, 5, 2, 6, 8, 3, 5, 19, 12, 4))
}