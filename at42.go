package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var s string
	fmt.Scan(&s)
	uCount := 0
	lCount := 0
	var result string
	for _, v := range s {
		if unicode.IsUpper(v) {
			uCount++
		} else {
			lCount++
		}
	}
	if uCount > lCount {
		result = strings.ToUpper(s)
	} else {
		result = strings.ToLower(s)
	}
	fmt.Println(result)

}
