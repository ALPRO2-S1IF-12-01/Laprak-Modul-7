package main

import (
	"fmt"
)

func main() {
	var x, y int
	var berat [1000]float64
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah (x y): ")
	fmt.Scan(&x, &y)

	if x > 1000 || x <= 0 || y <= 0 {
		fmt.Println("Nilai x harus 1-1000 dan y > 0")
		return
	}
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}
	jumlahWadah := (x + y - 1) / y 
	totalWadah := make([]float64, jumlahWadah)
	totalBeratSemua := 0.0

	for i := 0; i < x; i++ {
		idx := i / y
		totalWadah[idx] += berat[i]
		totalBeratSemua += berat[i]
	}
	fmt.Println("Total berat tiap wadah:")
	for _, total := range totalWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()
	rataRata := totalBeratSemua / float64(jumlahWadah)
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRata)
}
