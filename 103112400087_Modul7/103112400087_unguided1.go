package main

import (
	"fmt"
)

func main() {
	var n int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	fmt.Print("Masukkan berat anak kelinci ke-1: ")
	fmt.Scan(&berat[0])
	min, max := berat[0], berat[0]

	for i := 1; i < n; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f kg\n", min)
	fmt.Printf("Berat terbesar: %.2f kg\n", max)
}
