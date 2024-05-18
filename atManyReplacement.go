package main

import (
	"fmt"
)

// bufio, io, strconv 等を調べるとよい
// 完成したコード
func main() {
	var n, q int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	fmt.Scan(&q)
	c := make([]string, q)
	d := make([]string, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&c[i])
		fmt.Scan(&d[i])
	}
	from := map[string]string{"a": "a", "b": "b", "c": "c", "d": "d", "e": "e", "f": "f", "g": "g", "h": "h", "i": "i", "j": "j", "k": "k", "l": "l", "n": "n", "m": "m", "o": "o",
		"p": "p", "q": "q", "r": "r", "s": "s", "t": "t", "u": "u", "v": "v", "w": "w", "x": "x", "y": "y", "z": "z"}
	for i := 0; i < q; i++ {
		for key, v := range from {
			if v == c[i] {
				from[key] = d[i]
			}
		}
	}

	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = from[string(s[i])]
	}
	for _, v := range result {
		fmt.Printf("%s", v)
	}
}
