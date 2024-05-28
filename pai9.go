package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	sum = sum * (n - 1)
	fmt.Println(sum)
}
