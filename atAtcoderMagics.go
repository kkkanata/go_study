package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	// バッファリングされた入力を使用するためにbufio.Readerを作成
	reader := bufio.NewReader(os.Stdin)
	// バッファリングされた出力を使用するためにbufio.Writerを作成
	writer := bufio.NewWriter(os.Stdout)
	// プログラムの終了時にバッファをフラッシュしてすべての出力を表示
	defer writer.Flush()

	// 最初にnの値を読み取る
	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	// 2次元スライスacを作成し、各スライスは3つの整数を保持
	ac := make([][]int, n)
	for i := 0; i < n; i++ {
		ac[i] = make([]int, 3) // 各スライスの初期化
		// 入力からaとcの値を読み取る
		fmt.Fscanf(reader, "%d %d\n", &ac[i][0], &ac[i][1])
		// インデックスを1ベースで格納
		ac[i][2] = i + 1
	}

	// cの値に基づいてacスライスを昇順にソート
	sort.Slice(ac, func(i, j int) bool {
		return ac[i][1] < ac[j][1]
	})

	// 条件を満たす要素を選択し、結果をresultスライスに格納
	max := -1
	var result []int
	for i := 0; i < n; i++ {
		// aの値が現在の最大値より大きい場合に選択
		if max < ac[i][0] {
			result = append(result, ac[i][2]) // インデックスをresultに追加
			max = ac[i][0]                    // 現在の最大値を更新
		}
	}

	// 結果のインデックスを昇順にソート
	sort.Ints(result)

	// 結果の数を出力
	fmt.Fprintf(writer, "%d\n", len(result))
	// 各インデックスを出力
	for i, v := range result {
		if i == len(result)-1 {
			// 最後のインデックスの場合は改行を追加
			fmt.Fprintf(writer, "%d\n", v)
		} else {
			// それ以外の場合はスペースで区切る
			fmt.Fprintf(writer, "%d ", v)
		}
	}
}
