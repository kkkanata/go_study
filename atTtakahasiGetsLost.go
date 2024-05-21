package main

import (
	"fmt"
)

func main() {
	//入力処理
	var h, w, n int
	var t string
	fmt.Scan(&h, &w, &n)
	fmt.Scan(&t)
	s := make([]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}
	//探索処理
	tMap := make(map[byte]int, 4)
	for i := 0; i < n; i++ {
		tMap[t[i]]++
	}

	yInc := tMap['R'] - tMap['L']
	xInc := tMap['D'] - tMap['U']
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] != '#' && s[i+xInc][j+yInc] != '#' {
				for _, v := range t {
					var iw int
					var jw int
					switch v {
					case 'L':
						iw = 0
						jw -= 1
					case 'R':

					}
				}
			} else {
				continue
			}
		}
	}

}
