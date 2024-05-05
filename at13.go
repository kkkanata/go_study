package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	m := make([]int, len(s))
	less := 0
	for i := 0; i < len(s); {
		for j := less; j < len(t); j++ {
			if s[i] == t[j] {
				less = j + 1
				m[i] = j + 1
				i++
				break
			}
		}
	}
	for _, value := range m {
		fmt.Printf("%d ", value)
	}
}
