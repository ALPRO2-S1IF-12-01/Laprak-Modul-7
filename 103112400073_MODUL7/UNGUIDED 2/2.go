//Muhammad Zaky Mubarok
package main

import (
	"fmt"
)

func hitungTotalPerWadah(daftarBerat []float64, jumlahWadah int) []float64 {
	totalPerWadah := make([]float64, jumlahWadah)
	for i, berat := range daftarBerat {
		indeksWadah := i % jumlahWadah
		totalPerWadah[indeksWadah] += berat
	}
	return totalPerWadah
}

func hitungRataRata(daftar []float64) float64 {
	var total float64
	for _, nilai := range daftar {
		total += nilai
	}
	return total / float64(len(daftar))
}

func main() {
	var jumlahIkan, jumlahWadah int
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scan(&jumlahIkan)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&jumlahWadah)

	daftarBerat := make([]float64, jumlahIkan)
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&daftarBerat[i])
	}

	totalBeratPerWadah := hitungTotalPerWadah(daftarBerat, jumlahWadah)
	rataRataBeratWadah := hitungRataRata(totalBeratPerWadah)

	fmt.Println("Total berat ikan per wadah:")
	for i, berat := range totalBeratPerWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, berat)
	}

	fmt.Printf("Rata-rata berat ikan per wadah: %.2f kg\n", rataRataBeratWadah)
}
