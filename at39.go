package main

import (
	"fmt"
)

func main() {
	var n, l, r int
	fmt.Scan(&n)
	fmt.Scan(&l)
	fmt.Scan(&r)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}
	for i, j := l-1, r-1; i < j; i, j = i+1, j-1 {
		w := a[i]
		a[i] = a[j]
		a[j] = w
	}
	for i := 0; i < n; i++ {
		fmt.Print(a[i])
		fmt.Print(" ")
	}
}
