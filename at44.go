package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	fmt.Scan(&a)
	t := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
		if t[i] > sum {
			sum += (t[i] - sum) + a
		} else {
			sum += a
		}
		fmt.Println(sum)
	}

}
