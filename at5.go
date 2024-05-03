package main

import (
	"fmt"
)

func main() {
	var n int
	count := 0
	fmt.Scan(&n)
	var a, b string
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		switch a {
		case "G":
			if b == "C" {
				count++
			}
		case "C":
			if b == "P" {
				count++
			}
		case "P":
			if b == "G" {
				count++
			}
		}
	}
	fmt.Println(count)
}
