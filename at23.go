package main

import (
	"fmt"
	"strings"
)

func isValidAirportCode(s, t string) bool {
	// 空港コードの長さが3未満の場合は無効
	if len(s) < 3 {
		return false
	}

	if t[2] == 'X' {
		// t が X 付きの場合
		return (strings.Contains(s, string(t[0])) && strings.Contains(s, string(t[1])))
	}

	// t が X なしの場合
	for _, char := range t {
		if !strings.Contains(s, string(char)) {
			return false
		}
	}
	return true
}

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	if isValidAirportCode(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
