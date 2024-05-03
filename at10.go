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
		fmt.Scan(&a[i])
		fmt.Scan(&b[i])
	}
	sum := 0
	sum = b[0] - a[0] + 1
	max := b[0]
	less := a[0]
	for i := 1; i < n; i++ {
		diff1 := max - a[i]
		if max <= 
	}
	fmt.Printf("%d", sum)

}
