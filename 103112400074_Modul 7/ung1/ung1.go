//Ahmad Ruba'i
//103112400074 
package main

import "fmt"

const NMAX int = 1000

func main() {
	var n int
	var berat [NMAX]float64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf("%.2f %.2f\n", min, max)
}