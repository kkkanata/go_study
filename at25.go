package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var w int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var rs [][]int
	var b bool
	count := 0
	for !b {
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				rs = append(rs, []int{a[i+1], a[i]})
				w = a[i]
				a[i] = a[i+1]
				a[i+1] = w
				count++
			}

			b = true
			for i := 0; i < n-1; i++ {
				if a[i] > a[i+1] {
					b = false
				}
			}
		}
	}
	fmt.Println(count)
	for _, v := range rs {
		fmt.Printf("%d %d\n", v[0], v[1])
	}
}
