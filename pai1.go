package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}
	chocoCount, continuous := 0, 0
	bp := 0
	for i := 1; i <= n; i++ {
		if bp < m {
			if i == b[bp] {
				if bp < m {
					continuous++
					bp++
				}
			} else {
				continuous = 0
				chocoCount++
			}
		} else {
			continuous = 0
			chocoCount++
		}
		if continuous == k {
			fmt.Printf("%d", chocoCount)
			return
		}
	}
	fmt.Printf("%d", chocoCount)
}
