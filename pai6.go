package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	var x, y int
	fmt.Scan(&n)
	n2 := strings.Split(n,":")
	x, _ = strconv.Atoi(string(n2[0]))
	y, _ = strconv.Atoi(string(n2[1]))
	x -= 8
	if x < 0 {
		x += 24
	}
	fmt.Printf("%d:%d",x, y)


	
}
