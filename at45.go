package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 任意のスライスがM個の要素を持っているか確認
func isValidSlice(slice []string, M int) bool {
	return len(slice) == M
}

// 必要な最小のスライス数を求める関数
func minSlicesToCoverAllTrue(slices [][]string, M int) int {
	N := len(slices)
	needed := (1 << M) - 1 // 全ての要素がoになるためのビットマスク
	minCount := math.MaxInt32

	// ビットマスクを使用して全ての部分集合をチェック
	for subset := 1; subset < (1 << N); subset++ {
		combined := 0

		// 現在の部分集合に含まれるスライスを追加
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

		// すべての要素がoであるか確認
		if combined == needed {
			if count < minCount {
				minCount = count
			}
		}
	}

	if minCount == math.MaxInt32 {
		return -1 // 可能な組み合わせがない場合
	}
	return minCount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 最初の行を読み込み、NとMを取得
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

	// 入力データの検証
	for _, slice := range slices {
		if !isValidSlice(slice, M) {
			fmt.Println("不正なスライスがあります")
			return
		}
	}

	// 最小のスライス数を計算
	result := minSlicesToCoverAllTrue(slices, M)
	if result == -1 {
		fmt.Println("すべての要素をカバーすることができるスライスの組み合わせは存在しません")
	} else {
		fmt.Printf("%d", result)
	}
}
