package main

//C問題初見自力で解いた際のコード
import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	t := make([]float64, n+1)
	x := make([]float64, n+1)
	y := make([]float64, n+1)
	var md float64
	var tl float64
	t[0] = 0
	x[0] = 0
	y[0] = 0
	for i := 1; i < n+1; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}
	for i := 1; i < n+1; i++ {
		md = math.Abs(float64(x[i-1]-x[i])) + math.Abs(float64(y[i-1]-y[i]))
		tl = t[i] - t[i-1]
		if tl < md || int(tl)%2 != int(md)%2 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
