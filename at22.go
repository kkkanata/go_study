package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	var b bool
	fmt.Scan(&s, &t)
	s = strings.ToUpper(s)

	if t[2] == 'X' {
		for _, v1 := range t[:2] {
			b = false
			for _, v2 := range s {
				if v1 == v2 {
					b = true
					break
				}
			}
			if !b {
				fmt.Println("No")
				return
			}
		}
		fmt.Println("Yes")
	} else {
		for _, v1 := range t {
			b = false
			for _, v2 := range s {
				if v1 == v2 {
					b = true
					break
				}
			}
			if !b {
				fmt.Println("No")
				return
			}
		}
		fmt.Println("Yes")
	}
		
}	