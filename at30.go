package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	c := make([]int, n)
	acMin := make(map[int]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i], &c[i])
	}
	for i:=0;i<n;i++{ //Ciの大きさが10の9乗であるため0からcの最大値までの値を全て持つmapを作成すると10の9乗の要素を持つmapが必要
		acMin[c[i]] = -1
	}
	for i:=0;i<n;i++{
		if a[i] < acMin[c[i]] || acMin[c[i]] == -1 {
			acMin[c[i]] = a[i]
		}
	}
	max := -1
	for _, v := range c {
		if max < acMin[v] {
			max = acMin[v]
		}
	}
	fmt.Println(max)
}
