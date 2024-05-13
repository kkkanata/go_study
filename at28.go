package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	ri := -1
	first := h[0]
	for i := 1; i < n; i++ {
		if first < h[i] {
			ri = i + 1
			break
		}
	}
	fmt.Println(ri)
}
