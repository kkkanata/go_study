package main

import (
	"fmt"
	"strings"
)

func main() {
	var patterns []string = []string{"dream", "dreamer", "erase", "eraser"}
	var s string
	fmt.Scan(&s)
	var can bool
	for len(s) > 0 {
		can = false
		for _, value := range patterns {
			if strings.HasSuffix(s, value) {
				can = true
				s = s[:len(s)-len(value)]
				break
			}
		}
		if !can {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
