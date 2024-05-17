package main

import (
	"fmt"
)

//bufio, io, strconv 等を調べるとよい

func main() {
	var n, q int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)
	fmt.Scan(&q)
	sc := make(map[string][]int, n)
	for i := 0; i < n; i++ {
		sc[string(s[i])] = append(sc[string(s[i])], i)
	}
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = string(s[i])
	}
	var c, d string
	for i := 0; i < q; i++ {
		fmt.Scan(&c)
		fmt.Scan(&d)
		if _, ok := sc[c]; ok && c != d {
			for j := 0; j < len(sc[c]); j++ {
				result[sc[c][j]] = d
				//mapの更新処理
				sc[d] = append(sc[d], sc[c][j])
			}
			fmt.Println(result)
			delete(sc, c)
		}
	}
	for _, v := range result {
		fmt.Printf("%s", v)
	}
}
