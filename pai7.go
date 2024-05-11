package main

import (
	"fmt"
)

func main() {
	var x, y, p float64
	var s [][]float64
	for i := 0; i < 2; i++ {
		fmt.Scan(&x)
		fmt.Scan(&y)
		fmt.Scan(&p)
		s = append(s, []float64{x, y, p})
	}
	ws := make([]float64, 2)
	for i := 0; i < 2; i++ {
		ws[i] = s[i][2] / (s[i][0] * s[i][1])
	}
	if ws[0] == ws[1] {
		fmt.Println("DRAW")
	} else if ws[0] < ws[1] {
		fmt.Printf("%d %d %d\n", int(s[0][0]), int(s[0][1]), int(s[0][2]))
	} else {
		fmt.Printf("%d %d %d ", int(s[1][0]), int(s[1][1]), int(s[1][2]))
	}

}
