package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var a int
	count := 0
	for i:=0;i<n;i++{
		fmt.Scan(&a)
		count+=a
		if count < 0 {
			count = 0
		}
	}
	fmt.Println(count)


}