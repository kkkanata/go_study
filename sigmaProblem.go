package main

//5/14　もう一度覚える
import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	sum *= n - 1
	sort.Ints(a) //必要
	r := n
	cnt := 0
	for i := 0; i < n; i++ {
		r = max(r, i+1)
		for r-1 > i && a[r-1]+a[i] >= 100000000 {
			r--
		}
		cnt += n - r
	}
	res := sum - (100000000 * cnt)
	fmt.Println(res)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
