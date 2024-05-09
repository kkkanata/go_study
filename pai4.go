package main

import (
	"fmt"
)

func main() {
	var N, D int
	fmt.Scan(&N, &D)
	d := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		fmt.Scan(&d[i])
	}
	sum := D
	for _, v := range d {
		sum += D
		sum -= v
	}
	fmt.Printf("%d", sum*D)
}
