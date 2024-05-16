package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//chatGPTによる解説コピーコード 解析推奨
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Reading first list A
	reader.ReadString('\n')
	A := readInts(reader)

	// Reading second list B
	reader.ReadString('\n')
	B := readInts(reader)

	// Reading third list C
	reader.ReadString('\n')
	C := readInts(reader)

	// Create a set of sums
	S := make(map[int]struct{})
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				S[a+b+c] = struct{}{}
			}
		}
	}

	// Reading list X
	reader.ReadString('\n')
	X := readInts(reader)

	// Check if each x in X is in set S
	for _, x := range X {
		if _, found := S[x]; found {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func readInts(reader *bufio.Reader) []int {
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	strs := strings.Split(line, " ")
	nums := make([]int, len(strs))
	for i, s := range strs {
		nums[i], _ = strconv.Atoi(s)
	}
	return nums
}
