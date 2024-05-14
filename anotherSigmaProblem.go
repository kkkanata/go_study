package main

import (
	"fmt"
	"strconv"
)
//解説コードをchatGPTに変換させたコード
const MOD int = 998244353

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	d := make([]int, 11)
	for _, v := range a {
		d[len(strconv.Itoa(v))]++
	}

	res := 0
	p10 := make([]int, 11)
	p10[0] = 1
	for i := 1; i <= 10; i++ {
		p10[i] = p10[i-1] * 10 % MOD
	}

	for i, v := range a {
		res = (res + v*i) % MOD
		d[len(strconv.Itoa(v))]--
		for j := 1; j <= 10; j++ {
			res = (res + p10[j]*v*d[j]) % MOD
		}
	}

	fmt.Println(res)
}
