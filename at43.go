package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	if s == "AtCoder" && t == "Land" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
