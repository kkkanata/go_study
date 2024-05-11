package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Printf("%.1f\n", math.Pi)
	case 2:
		fmt.Printf("%.2f\n", math.Pi)
	case 3:
		fmt.Printf("%.3f\n", math.Pi)
	case 4:
		fmt.Printf("%.4f\n", math.Pi)
	case 5:
		fmt.Printf("%.5f\n", math.Pi)
	case 6:
		fmt.Printf("%.6f\n", math.Pi)
	case 7:
		fmt.Printf("%.7f\n", math.Pi)
	case 8:
		fmt.Printf("%.8f\n", math.Pi)
	case 9:
		fmt.Printf("%.9f\n", math.Pi)
	case 10:
		fmt.Printf("%.10f\n", math.Pi)
	case 11:
		fmt.Printf("%.11f\n", math.Pi)
	case 12:
		fmt.Printf("%.12f\n", math.Pi)
	case 13:
		fmt.Printf("%.13f\n", math.Pi)
	case 14:
		fmt.Printf("%.14f\n", math.Pi)
	case 15:
		fmt.Printf("%.15f\n", math.Pi)
	}

}
