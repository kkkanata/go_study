package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ct := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ct += ((a[i] + a[j]) % 100000000)
		}
	}
	fmt.Println(ct)
}
