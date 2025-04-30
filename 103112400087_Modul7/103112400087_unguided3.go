package main

import (
	"fmt"
)

func main() {
	var n int
	var berat [100]float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	total := 0.0
	min, max := 0.0, 0.0

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
		total += berat[i]

		if i == 0 || berat[i] < min {
			min = berat[i]
		}
		if i == 0 || berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", total/float64(n))
}
