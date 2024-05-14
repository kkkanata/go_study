package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	n := len(s)
	duplicate := make(map[rune]int, n)
	for _, v := range s {
		duplicate[v]++
	}
	dupSum := 0
	sum := (n * (n - 1)) / 2 //最大swap回数を求める
	for _, v := range duplicate {
		dupSum += v - 1
	}
	dupSum = ((dupSum + 1) * (dupSum)) / 2
	if sum-dupSum == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(sum - dupSum)
	}
}
