package main

import (
	"fmt"
)

//不完全 入れ替えの順序の規則性に着目

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	var w int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var rs [][]int
	count := 0
	min := 100000
	k, w, wn := 0, 0, 0
	var b, b2 bool
LOOP:
	for k < n-1 {
		b = false
		min = 100000
		for i := k; i < n; i++ {
			if min > a[i] {
				min = a[i]
				wn = i
				b = true
			}
		}
		if b {
			rs = append(rs, []int{min, a[k]})
			w = a[k]
			a[k] = min
			a[wn] = w
			count++
		}
		b2 = true
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				b2 = false
			}
		}
		if b2 {
			break LOOP
		}
		k++
	}
	fmt.Println(count)
	for _, v := range rs {
		fmt.Printf("%d %d\n", v[0], v[1])
	}
}
