package main

import (
	"fmt"
)

func main() {
	n := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scan(&n[i])
	}
	min := 1001
	max := 0
	for i := 0; i < 3; i++ {
		if n[i] < min {
			min = n[i]
		}
		if n[i] > max {
			max = n[i]
		}
	}
	var medium int
	for i := 0; i < 3; i++ {
		if n[i] > min && n[i] < max {
			medium = n[i]
		}
	}

	fmt.Println(medium)
}
