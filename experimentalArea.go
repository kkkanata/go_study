package main

import (
	"fmt"
)

func main() {
	var s string
	s = "123"
	fmt.Println(string(int(s[2] + 1)))
}
