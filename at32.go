package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	res := n * n
	cnt := make(map[rune]int, n)
	for _, v := range s {
		cnt[v]++
	}
	var b bool
	for _, v := range cnt {
		res -= v * v
		if v > 1 {
			b = true
		}
	}
	res /= 2
	if b {
		res++
	}
	fmt.Println(res)
}
