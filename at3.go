package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s)+2; i++ {
		fmt.Printf("+")
	}
	fmt.Printf("\n+%s+\n", s)
	for i := 0; i < len(s)+2; i++ {
		fmt.Printf("+")
	}

}
