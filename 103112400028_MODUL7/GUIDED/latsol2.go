// MUHHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	var beratIkan [1000]float64
	fmt.Println("\nMasukkan berat setiap ikan (dalam kg):")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	var totalBeratWadah []float64
	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += beratIkan[j]
		}
		totalBeratWadah = append(totalBeratWadah, total)
	}

	// Output total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total := range totalBeratWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, total)
	}

	// Hitung rata-rata berat ikan di setiap wadah
	totalBerat := 0.0
	for _, berat := range totalBeratWadah {
		totalBerat += berat
	}
	rataRata := totalBerat / float64(len(totalBeratWadah))

	// Output rata-rata berat ikan di setiap wadah
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f kg\n", rataRata)
}
