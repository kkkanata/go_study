package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	S := make([]bool, N)
	var w int
	for i := 0; i < N; i++ {
		fmt.Scan(&w)
		if w == 1 {
			S[i] = true
		}
	}
	var l, r int
	for i := 0; i < N; i++ {
		if S[i] {

		}
	}

}
