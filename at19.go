package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	maxLength := 1
	smax := make([]int, 8)
	smax2 := make([]int, 8)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//横の探索+
			for inc := 1; i+inc < n && int(s[i+inc][j]) == int(s[i][j])+inc; inc++ {
				smax[0] = inc + 1
			}
			if smax2[0] < smax[0] {
				smax2[0] = smax[0]
			}
			//横の探索-
			for inc := 1; i+inc < n && int(s[i+inc][j]) == int(s[i][j])-inc; inc++ {
				smax[1] = inc + 1
			}
			if smax2[1] < smax[1] {
				smax2[1] = smax[1]
			}
			//縦の探索+
			for inc := 1; j+inc < n && int(s[i][j+inc]) == int(s[i][j])+inc; inc++ {
				smax[2] = inc + 1
			}
			if smax2[2] < smax[2] {
				smax2[2] = smax[2]
			}
			//縦の探索-
			for inc := 1; j+inc < n && int(s[i][j+inc]) == int(s[i][j])-inc; inc++ {
				smax[3] = inc + 1
			}
			if smax2[3] < smax[3] {
				smax2[3] = smax[3]
			}
			//左斜めの探索+
			for incx, incy := 1, -1; i+incx < n && j+incy >= 0 && int(s[i+incx][j+incy]) == int(s[i][j])+incx; incx, incy = incx+1, incy-1 {
				smax[4] = incx + 1
			}
			if smax2[4] < smax[4] {
				smax2[4] = smax[4]
			}
			//左斜めの探索-
			for incx, incy := 1, -1; i+incx < n && j+incy >= 0 && int(s[i+incx][j+incy]) == int(s[i][j])-incx; incx, incy = incx+1, incy-1 {
				smax[5] = incx + 1
			}
			if smax2[5] < smax[5] {
				smax2[5] = smax[5]
			}
			//右斜めの探索+
			for incx, incy := 1, 1; i+incx < n && j+incy < n && int(s[i+incx][j+incy]) == int(s[i][j])+incx; incx, incy = incx+1, incy+1 {
				smax[6] = incx + 1
			}
			if smax2[6] < smax[6] {
				smax2[6] = smax[6]
			}
			//右斜めの探索-
			for incx, incy := 1, 1; i+incx < n && j+incy < n && int(s[i+incx][j+incy]) == int(s[i][j])-incx; incx, incy = incx+1, incy+1 {
				smax[7] = incx + 1
			}
			if smax2[7] < smax[7] {
				smax2[7] = smax[7]
			}

		}
	}
	for _, value := range smax2 {
		if maxLength < value {
			maxLength = value
		}
	}
	fmt.Printf("%d", maxLength)
}
