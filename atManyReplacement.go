package main

import (
	
	"fmt"
)
//bufio, io, strconv 等を調べるとよい

func main() {
	var n, q int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	fmt.Scan(&q)
	result := make([]string, n)
	for i:=0;i<n;i++{
		result[i] = string(s[i])
	}
	var c, d string
	for i:=0;i<q;i++{
		fmt.Scan(&c)
		fmt.Scan(&d)
		for j:=0; j<n;j++{
			if string(result[j]) == c {	
			result[j] = d
			}
		}
	}
	for _, v := range result {
		fmt.Printf("%s", v)
	}
}
