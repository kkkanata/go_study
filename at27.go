package main

import (
	"fmt"
	"math"
)
//不完全なコード
func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var rs [][]int
	min := math.MaxInt32
	var minNum int
	//var relative int
	var b bool
	var rb bool
	count := 0
	k := 0;
	for ; k < n-1 && !rb; k++ {
		for i:=k; i < n; i++ {
			if min > a[i] {
				min = a[i]
				minNum = i
				b=true
			}
		}
		if b {
			rs = append(rs, []int{k+1, minNum+1})
			w := a[k]
			a[minNum] = a[k]
			a[k] = w
			count++
		}
		rb = true
		for i:=0; i<n-1; i++ {
			if a[i] > a[i+1] {
				rb = false
				break
			}
		}
	}
	fmt.Println(count)
	for _, v := range rs {
		fmt.Printf("%d %d\n",v[0],v[1])
	}
}
