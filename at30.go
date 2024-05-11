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
		ct += a[i] * (n - i - 1) * a[i+1]
		ct %= 100000000
	}
	fmt.Println(ct)
}
