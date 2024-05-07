package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&a[i])
	}
	var sum int
	sum = 0
	for _, v := range a {
		if v <= m {
			sum += v
		}
	}
	fmt.Println(sum)
}
