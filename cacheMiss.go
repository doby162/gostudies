package main

import (
	"fmt"
	"time"
)

const size = 20000

var final [size][size]byte

func main() {
	var t = time.Now()
	var arr [size][size]byte

	t = time.Now()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr[j][i] = byte(i + j)
		}
	}
	final = arr
	fmt.Println(time.Now().Sub(t))

	t = time.Now()
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			arr[i][j] = byte(i + j)
		}
	}
	final = arr
	fmt.Println(time.Now().Sub(t))
}
