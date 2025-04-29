// M.HANIF AL FAIZ
// 103112400042
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
	wadahCount := (x + y - 1) / y

	for i := 0; i < wadahCount; i++ {
		total := 0.0
		start := i * y
		end := start + y
		if end > x {
			end = x
		}
		for j := start; j < end; j++ {
			total += beratIkan[j]
		}
		totalBeratWadah = append(totalBeratWadah, total)
	}

	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, total := range totalBeratWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, total)
	}

	totalBerat := 0.0
	for _, berat := range totalBeratWadah {
		totalBerat += berat
	}
	rataRata := totalBerat / float64(len(totalBeratWadah))

	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f kg\n", rataRata)
}
