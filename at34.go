package main

import (
	"fmt"
)
//要件は達成したがbufio, io, strconv 等を使用するコードと比較すると数倍低速、そのあたりを調べるとよい
func main() {
	var n, m, l, q int
	fmt.Scan(&n)
	A:= make([]int, n)
	for i:= 0;i<n;i++{
		fmt.Scan(&A[i])
	}
	fmt.Scan(&m)
	B := make([]int, m)
	for i:= 0;i<m;i++{
		fmt.Scan(&B[i])
	}
	fmt.Scan(&l)
	C := make([]int, l)
	for i:= 0;i<l;i++{
		fmt.Scan(&C[i])
	}
	fmt.Scan(&q)
	X := make([]int, q)
	for i:= 0;i<q;i++{
		fmt.Scan(&X[i])
	}


	S := make(map[int]struct{})
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				S[a+b+c] = struct{}{}
			}	
		}
	}
	for _, x := range X {
		if _, found := S[x]; found {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
