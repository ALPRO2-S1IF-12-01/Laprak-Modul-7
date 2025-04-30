// Anastasia Adinda Narendra Indrianto
//103112400085
package main

import (
	"fmt"
)

func main() {
	var anastasiaX, y int
	fmt.Print("Masukkan jumlah ikan (anastasiaX) dan kapasitas wadah (y): ")
	fmt.Scan(&anastasiaX, &y)

	if anastasiaX <= 0 || anastasiaX > 1000 || y <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	var ikan [1000]float64
	fmt.Println("Masukkan berat ikan satu per-satu:")
	for i := 0; i < anastasiaX; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (anastasiaX + y - 1) / y // pembulatan ke atas untuk wadah terakhir
	var totalWadah [1000]float64
	var totalIkanInWadah [1000]int // untuk menghitung jumlah ikan dalam setiap wadah

	for i := 0; i < anastasiaX; i++ {
		idx := i / y
		totalWadah[idx] += ikan[i]
		totalIkanInWadah[idx]++
	}

	var totalBerat float64
	for i := 0; i < jumlahWadah; i++ {
		if totalIkanInWadah[i] > 0 {
			// Menghitung berat rata-rata ikan dalam wadah
			rataRataIkan := totalWadah[i] / float64(totalIkanInWadah[i])
			fmt.Printf("Berat rata-rata ikan dalam wadah %d: %.2f\n", i+1, rataRataIkan)
		} else {
			fmt.Println("Tidak ada ikan dalam wadah ini.")
		}
		totalBerat += totalWadah[i]
	}

	fmt.Printf("\nTotal berat semua ikan: %.2f\n", totalBerat)
}
