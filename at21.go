package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scan(&n)
	count := 0
	for i := 0; i < 4; i++ {
		count = 0
		for _, v := range n {
			if rune(n[i]) == v {
				count++
			}
			if count >= 2 {
				fmt.Println("NG")
				return
			}
		}
	}
	fmt.Println("OK")
}