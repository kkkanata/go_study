package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}
	maxNum := 0
	sum := 0
	for i := 0; i < n; i++ {
		if b[maxNum]-a[maxNum] < b[i]-a[i] {
			maxNum = i
		}
		sum += a[i]
	}
	fmt.Printf("%d", sum+b[maxNum]-a[maxNum])
}
