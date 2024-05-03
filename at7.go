package main

import (
	"fmt"
)

func main() {
	var n, l, x int
	fmt.Scan(&n, &l)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if l > x {
			l += x / 2
		} else if l < x {
			l /= 2
		}
	}
	fmt.Printf("%d", l)
}
