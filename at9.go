package main

import (
	"fmt"
)

// 不可
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
		diffab := a[i] - b[i-1]
		if diffab <= 1 {
			diffa := a[i-1] - a[i]
			diffb := b[i] - b[i-1]
			if less >= a[i] && max <= b[i] {
				sum = sum + diffa + diffb
			} else if less >= a[i] {
				sum = sum + diffa
			} else if max <= b[i] {
				sum = sum + diffb
			}
		} else if sum < b[i]-a[i] {
			sum = b[i] - a[i] + 1
		}

	}
	fmt.Printf("%d", sum)

}
