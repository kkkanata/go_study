package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &s, &q)

	// 文字列をruneスライスに変換
	resultSlice := []rune(s)

	// スキャナーを初期化
	scanner := bufio.NewScanner(os.Stdin)

	// 置換用のマップ
	replacementMap := make(map[rune]rune)

	// 置換情報を取得
	for i := 0; i < q; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			if len(line) == 3 && line[1] == ' ' {
				c := rune(line[0])
				d := rune(line[2])
				replacementMap[c] = d
			}
		}
	}

	// 文字列の置換処理
	for i, v := range resultSlice {
		if newChar, ok := replacementMap[v]; ok {
			resultSlice[i] = newChar
		}
	}

	// 結果を出力
	fmt.Println(string(resultSlice))
}
