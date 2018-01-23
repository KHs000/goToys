package main

import "fmt"

func getGreater (xs ...int) (greatest int) {
	greatest = -1
	for _, x := range xs {
		if x >= greatest {greatest = x}
	}

	return
}

func fib (n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fib (n - 1) + fib (n - 2)
	}
}

func square(x *float64) float64{
  	*x = *x * *x
  	return *x
}

func swapIntegers (x, y *int) (newX, newY int) {
	swap := *x
	newX = *y
	newY = swap

	return
}

func main () {
  	fmt.Println(fib(5))
}