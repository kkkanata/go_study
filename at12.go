package main

import (
	"fmt"
)

func main() {
	var n, x, y, z int
	fmt.Scan(&n, &x, &y, &z)
	if z >= x && z <= y || z <= x && z >= y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
