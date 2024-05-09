package main

import (
	"fmt"
	"strings"
)
 //at20.go 正当

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	count := 0
	var b bool
	s = strings.ToUpper(s)
	
	for i := 0; i < len(t); {
		b = false
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				s = s[j+1:]
				b = true
				count++
				i++
				break
			}
		}
		if !b {
			break
		}
	}
	if count >= 3 {
		fmt.Println("Yes")
	} else if t[2] == 'X' && count >= 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
