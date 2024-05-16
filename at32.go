package main

import (
	"fmt"
)
//コピーコード よく見なければならない

func main() {
	var s string
	var same bool
	fmt.Scan(&s)
	n := len(s)
	cnt := make(map[rune]int, n)
	for _, v := range s {
		cnt[v]++
	}
	res:=(n*(n-1)) //n*nにしてv*vを引いても同じ
	for _, v := range cnt {
		res -=(v*(v-1)) //n*nにしてv*vを引いても同じ
		if v>1 {same=true}
	}
	res /= 2
	if same {res++}
	fmt.Println(res)
}

