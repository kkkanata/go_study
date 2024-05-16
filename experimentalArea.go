package main

import (
	"fmt"
)

func main() {
	var n, ans int64
	var s string

	fmt.Scan(&s)
	n = int64(len(s))
	cnt := make([]int64, 26)
	same := false

	for i := 0; i < int(n); i++ {
		cnt[s[i]-'a']++
	}
	ans = n * n
	for i := 0; i < 26; i++ {
		ans -= cnt[i] * cnt[i]
		if cnt[i] > 1 {
			same = true
		}
	}
	ans /= 2
	if same {
		ans++
	}
	fmt.Println(ans)
}
