package main

import (
	"fmt"
)
//一応完成したコード
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
	}
	m := make(map[int]bool, n)

	for _, v := range a {
		if v >= 1 && v <= k {
			m[v] = true
		}
	}
	dec := 0
	for _, v := range a{ //ここで重複している
		if m[v] {
			dec+= v
			m[v] = false
		} 
	}
	sum := ((1+k)*k)/2
	//fmt.Println(sum)
	fmt.Println(sum - dec)
}

