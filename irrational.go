package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(Sqrt(9, 10))
}

func Sqrt(x float64, i int) float64 {
	z := 1.0 //our initital guess
	for ; i > 0; i-- {
		z = z - (z*z-x)/(2*z)
	}
	return z
}
