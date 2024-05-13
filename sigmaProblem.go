package main

func main() {
	var n int
	fmt.Scan(&n)
	a:=make([]int, n)
	sum := 0
	for i:=0;i<n;i++{
		fmt.Scan(&a[i])
		sum+=a[i]
	}
	sum = sum/(n-1)

}
