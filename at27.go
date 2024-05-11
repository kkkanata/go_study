package main

import (
	"fmt"
)

// 完成したコード見直すべき
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		pos[a[i]-1] = i
	}
	var w, w2 int
	var rs [][]int
	count := 0
	for i := 0; i < n-1; i++ {
		if i+1 != a[i] {
			rs = append(rs, []int{a[pos[i]], pos[i] + 1})
			w = a[pos[i]]
			w2 = pos[i]
			a[pos[i]] = a[i]
			pos[i] = pos[a[i]-1]
			pos[a[i]-1] = w2
			a[i] = w
			count++
		}
	}
	fmt.Println(count)
	for _, v := range rs {
		fmt.Printf("%d %d\n", v[0], v[1])
	}
}
