package main

import (
	"fmt"
)
//自分なりに完成させたコード
func main() {
	//入力処理
	var h, w, n int
	var t string
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&n)
	fmt.Scan(&t)
	xinc, yinc := 0, 0
	for _, v := range t {
		switch v {
		case 'L':
			yinc -= 1
		case 'R':
			yinc += 1
		case 'U':
			xinc -= 1
		case 'D':
			xinc += 1
		default:
			fmt.Println("error")
		}
	}
	s := make([]string, h)
	for i:=0;i<h;i++{
		fmt.Scan(&s[i])
	}
	var xw, yw int
	//探索
	count := 0
	for x := 1; x < h - 1; x++{
		for y := 1; y < w - 1; y++{
			b := true
			if s[x][y] == '.' && x+xinc < h && x+xinc >= 0 && y+yinc < w && y+yinc >= 0 && s[x+xinc][y+yinc] == '.' {//問題
				xw, yw = x, y
				for _, v := range t {
					switch v {
					case 'L':
						if y-1 >= 0 && s[xw][yw-1] != '#' {
							yw -= 1
						} else {
							b = false
						}
					case 'R':
						if y+1 < w && s[xw][yw+1] != '#' {
							yw += 1
						} else {
							b = false
						}
					case 'U':
						if x - 1 >= 0 && s[xw-1][yw] != '#' {
							xw -= 1
						} else {
							b = false
						}
					case 'D':
						if x+1 < h && s[xw+1][yw] != '#' {
							xw += 1
						} else {
							b = false
						}
					}
				}
			} else {
				b = false
			}
			if b {
				count++
			}
		}
	}
	fmt.Println(count)
}
