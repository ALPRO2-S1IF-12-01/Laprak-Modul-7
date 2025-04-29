package main

import (
	"fmt"
)

func main() {
	var N int
	var berat [1000]float64
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)
	if N > 1000 || N <= 0 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}
	minBerat := berat[0]
	maxBerat := berat[0]
	for i := 1; i < N; i++ {
		if berat[i] < minBerat {
			minBerat = berat[i]
		}
		if berat[i] > maxBerat {
			maxBerat = berat[i]
		}
	}
	fmt.Printf("Berat terkecil: %.2f\n", minBerat)
	fmt.Printf("Berat terbesar: %.2f\n", maxBerat)
}
