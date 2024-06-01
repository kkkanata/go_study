package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	a := make([]int, m)
	xSum := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
		xSum[i] = 0
	}
	var x int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			xSum[j] += x
		}
	}
	for i := 0; i < m; i++ {
		if a[i] > xSum[i] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
