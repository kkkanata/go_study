package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力からnを取得
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter n: ")
	input, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(input[:len(input)-1])

	// グリッドのサイズ
	size := pow(3, n)

	// グリッドを初期化
	grid := make([][]rune, size)
	for i := range grid {
		grid[i] = make([]rune, size)
		for j := range grid[i] {
			grid[i][j] = '#'
		}
	}

	// グリッドを更新する関数
	updateGrid(grid, 0, 0, size)

	// グリッドを出力
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Println()
	}
}

// グリッドを更新する関数
func updateGrid(grid [][]rune, x, y, size int) {
	if size == 1 {
		return
	}

	subSize := size / 3

	// 中央のサブグリッド全体を'.'にする
	for i := x + subSize; i < x+2*subSize; i++ {
		for j := y + subSize; j < y+2*subSize; j++ {
			grid[i][j] = '.'
		}
	}

	// 各サブグリッドに対して再帰的に処理を行う
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i != 1 || j != 1 {
				updateGrid(grid, x+i*subSize, y+j*subSize, subSize)
			}
		}
	}
}

// 3のべき乗を計算する関数
func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
