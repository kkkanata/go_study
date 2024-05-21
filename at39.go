package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	count := 0
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
		count+=a[i]
		if count < 0 {
			count = 0
		}
	}
	fmt.Println(count)


}