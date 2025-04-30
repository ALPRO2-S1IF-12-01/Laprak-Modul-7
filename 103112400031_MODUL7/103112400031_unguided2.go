// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan banyak ikan yang akan dijual: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan kapasitas ikan per wadah: ")
	fmt.Scan(&y)

	var beratIkan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%v (kg): ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (x + y - 1) / y
	var totalWadah [1000]float64
	var totalBerat float64

	for i := 0; i < x; i++ {
		wadahKe := i / y
		totalWadah[wadahKe] += beratIkan[i]
	}

	fmt.Println("\nTotal berat ikan per wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %v: %.2f kg\n", i+1, totalWadah[i])
		totalBerat += totalWadah[i]
	}

	rataRata := totalBerat / float64(jumlahWadah)
	fmt.Printf("\nRata-rata berat ikan per wadah: %.2f kg\n", rataRata)
}