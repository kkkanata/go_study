package main

import (
	"fmt"
) // 自作コード(不完全)
func main() {
	var s string
	fmt.Scan(&s)
	var ss []string = []string{"dream", "dreamer", "erase", "eraser"}
	max := len(s)
OuterLoop:
	for max > 0 {
		for i := 0; i < len(ss) && max != 0; i++ {
			less := max - len(ss[i])
			if less >= 0 {
				if s[less:max] == ss[i] {
					max = max - len(ss[i])
					i = 0
				} else if i == len(ss)-1 {
					break OuterLoop
				}
			}
		}
	}
	if max == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
