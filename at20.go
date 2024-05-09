package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	count := 0

	
	if t[2] == 'X' {
		for i := 0; i < len(t)-1; i++ {
			for j := 0; j < len(s); j++ {
				if t[i]+32 == s[j] {
					s = s[:j] + s[j+1:]
					count++
					break
				}
			}
		}
		if count >= 2 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	} else { 
		for i := 0; i < len(t); i++ {
			for j := 0; j < len(s); j++ {
				if t[i]+32 == s[j] {
					s = s[:j] + s[j+1:]
					count++
					break
				}
			}
		}
		if count >= 3 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
