package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)
	for i := 1; i <= n; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println("AB")
		} else if i%x == 0 {
			fmt.Println("A")
		} else if i%y == 0 {
			fmt.Println("B")
		} else {
			fmt.Println("N")
		}
	}
}
