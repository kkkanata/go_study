package main
//メモ化再帰
import (
	
	"fmt"
)
//再帰処理を行う
/*func f(n int) int {
	if n == 1 {
		return 0
	} else {
		return f(n/2) + f((n+1)/2) + n
	}
}
*/

func recursion(n int) int {
	m := make(map[int]int)
	var f func(int)int
	f = func (n int) int {
		if n==1 {
		return 0
		}
		if _, ok := m[n]; ok {
			return m[n]	
		}
		m[n] = f(n/2) + f((n+1)/2) + n
		return m[n]
	}
	result := f(n)
	return result
}

func main() {
	var n int
	fmt.Scan(&n)
	result := recursion(n)
	fmt.Println(result)
}
