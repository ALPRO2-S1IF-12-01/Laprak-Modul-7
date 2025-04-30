// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0.")
		return
	}

	beratIkan := make([]float64, x)
	fmt.Println("\nMasukkan berat setiap ikan (dalam kg):")
	for i := 0; i < x; i++ {
		fmt.Print("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	wadahCount := (x + y - 1) / y
	totalBeratWadah := make([]float64, wadahCount)

	for i := 0; i < x; i++ {
		totalBeratWadah[i/y] += beratIkan[i]
	}

	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total := range totalBeratWadah {
		fmt.Printf("Wadah %d: %2.f kg\n", i+1, total)
	}

	totalBerat := 0.0
	for _, berat := range totalBeratWadah {
		totalBerat += berat
	}
	rataRata := totalBerat / float64(len(totalBeratWadah))

	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %2.f kg\n", rataRata)
}
