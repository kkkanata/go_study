package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a[i], _ = strconv.Atoi(scanner.Text())
	}
	ct := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ct += ((a[i] + a[j]) % 100000000)
		}
	}
	fmt.Println(ct)
}
