package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	ac := make([][]int, n)
	/*for i := 0; i < n; i++ {
		ac[i] = make([]int, 3)
	}*/

	//ac[i][0]=Aiの値ac[i][1]=Ciの値
	for i := 0; i < n; i++ {
		ac[i] = make([]int, 3)
		fmt.Scan(&ac[i][0])
		fmt.Scan(&ac[i][1])
		ac[i][2] = i + 1
	}
	sort.Slice(ac, func(i, j int) bool {
		return ac[i][1] < ac[j][1]
	})
	max := -1
	count := 0
	var result []int
	for i := 0; i < n; i++ {
		if max < ac[i][0] {
			result = append(result, ac[i][2])
			max = ac[i][0]
			count++
		}
	}
	sort.Ints(result)
	fmt.Println(count)
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
