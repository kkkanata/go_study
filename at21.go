package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    A := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&A[i])
    }

    less, upper := 100001, 0
    lessAndUpp := make([][]int, 0)

    for k := 0; ; k++ {
        lessNum, uppNum := -1, -1
        for i := k; i < n; i++ {
            if less > A[i] {
                less = A[i]
                lessNum = i
            }
        }
        for i := k; i < n; i++ {
            if upper < A[i] {
                upper = A[i]
                uppNum = i
            }
        }
        if lessNum == -1 || uppNum == -1 {
            break
        }
        lessAndUpp = append(lessAndUpp, []int{A[lessNum], A[uppNum]})
        A[lessNum], A[uppNum] = A[uppNum], A[lessNum]

        b := true
        for i := 1; i < n && b; i++ {
            if A[i-1] <= A[i] {
                b = false
            }
        }
        if b {
            fmt.Println(k + 1)
            for _, v := range lessAndUpp {
                fmt.Printf("%d %d\n", v[0], v[1])
            }
            break
        }
    }
}
