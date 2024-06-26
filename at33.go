package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	cnt := make(map[rune]int, n)
	for _, v := range s {
		cnt[v]++
	}
	cntSum := 0
	sum := (n * (n - 1)) / 2
	for _, v := range cnt {
		cntSum += v
	}
	cntSum = ((cntSum + 1) * (cntSum)) / 2
	if sum-dupSum == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(sum - dupSum)
	}
}