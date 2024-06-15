package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isValidSlice(slice []string, M int) bool {
	return len(slice) == M
}

func minSlicesToCoverAllTrue(slices [][]string, M int) int {
	N := len(slices)
	needed := (1 << M) - 1
	minCount := math.MaxInt32

	for subset := 1; subset < (1 << N); subset++ {
		combined := 0

		count := 0
		for i := 0; i < N; i++ {
			if subset&(1<<i) != 0 {
				count++
				for j := 0; j < M; j++ {
					if slices[i][j] == "o" {
						combined |= (1 << j)
					}
				}
			}
		}

		if combined == needed {
			if count < minCount {
				minCount = count
			}
		}
	}

	if minCount == math.MaxInt32 {
		return -1
	}
	return minCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Split(firstLine, " ")
	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	slices := make([][]string, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		line := scanner.Text()
		slices[i] = strings.Split(line, "")
	}

	for _, slice := range slices {
		if !isValidSlice(slice, M) {
			return
		}
	}

	result := minSlicesToCoverAllTrue(slices, M)
	if result == -1 {
	} else {
		fmt.Printf("%d", result)
	}
}
