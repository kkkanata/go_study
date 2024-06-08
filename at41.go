package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	count := 0
	for _, v := range h {
		if m >= v {
			m -= v
			count++
		} else {
			break
		}
	}
	fmt.Println(count)
}
