package main

import (
	"fmt"
)

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	max, _ := MaxAndMin(10, 20)
	fmt.Printf("max=%d\n", max)
}
