package main

import (
	"fmt"
)

func recursion(n int) int {
	if n == 1 {
		return 0
	} else {
		return recursion(n/2) + recursion((n+1)/2) + n
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	result := recursion(n)
	fmt.Println(result)
}
