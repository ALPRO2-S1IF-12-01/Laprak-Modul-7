// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	total := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		total[i/y] += berat[i]
	}
	for _, t := range total {
		fmt.Printf("%.2f ", t)
	}
	fmt.Println()
	var sum float64
	for _, t := range total {
		sum += t
	}
	fmt.Printf("%.2f\n", sum/float64(jumlahWadah))
}