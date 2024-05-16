package main

import (
	
	"fmt"
	"strconv"
)
//bufio, io, strconv 等を調べるとよい

func check(n int) bool {
	s := strconv.Itoa(n)
	r := []rune(s)
	for i, j := 0, len(s) -1 ; i<j; i, j = i+1, j-1 {
		w := r[i]
		r[i] = r[j]
		r[j] = w
	}
	return s == string(r)
}

func main() {
	var n int
	fmt.Scan(&n)
	max := 1
	for i:= 1; i*i*i <= n; i++ {
		if check(i*i*i) {max = i*i*i}
	}
	fmt.Println(max)
}
