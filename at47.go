package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		fmt.Scan(&a[i])
	}
	count := 0
	for i := 0; i < 2*n-2; i++ {
		if a[i] == a[i+2] {
			count++
		}
	}
	fmt.Println(count)
}
