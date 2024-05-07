package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum := 1
	for ;x != sum && x > sum; {
		sum *= 2
	}
	if x == sum {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}

}
