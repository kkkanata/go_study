package main

import (
	"fmt"
	"strings"
)
//不完全なコード
//部分列 tに対してsの文字が同じ進行方向で並べられている可能性

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	count := 0
	s = strings.ToUpper(s)
	
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				s = s[:j] + s[j+1:]
				count++
				break
			}
		}
	}
	if count == 3 {
		fmt.Println("Yes")
	} else if t[2] == 'X' && count == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
