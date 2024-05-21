package main
//メモ化再帰
import (
	"fmt"
)
//再帰処理を行う

var m = make(map[int]int)
//再帰をしていく中でfにたいして同じ値を渡すことが多くある場合がある
//一度計算した値をmapにメモしておくことで一度行った処理を省略できる
func f(n int) int {
	if n == 1 {
		return 0
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = f(n/2)+f((n+1)/2)+ n
	return m[n]
}


func main() {
	var n int
	fmt.Scan(&n)
	result := f(n)
	fmt.Println(result)
}