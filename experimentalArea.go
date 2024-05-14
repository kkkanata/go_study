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

	// 計算する必要があるのは、各文字が何回重複しているかを考慮すること
	// 具体的には、重複する文字のペア数を減らす
	dupSum := 0
	for _, v := range duplicate {
		if v > 1 {
			dupSum += (v * (v - 1)) / 2 // 重複する文字のペア数
		}
	}

	// 最大スワップ回数は文字列の長さ n の中での組み合わせ (n * (n - 1)) / 2
	sum := (n * (n - 1)) / 2

	// 実際のスワップ回数は、重複する文字のペア数を引いたもの
	actualSwaps := sum - dupSum

	if actualSwaps == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(actualSwaps)
	}
}
