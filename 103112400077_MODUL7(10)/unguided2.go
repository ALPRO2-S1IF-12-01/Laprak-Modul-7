package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan banyak ikan dan jumlah per wadah: ")
	fmt.Scan(&x, &y)

	var beratIkan [1000]float64
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	totalWadah := x / y
	if x%y != 0 {
		totalWadah++
	}

	var totalPerWadah [1000]float64
	idx := 0
	for i := 0; i < totalWadah; i++ {
		var total float64 = 0
		for j := 0; j < y && idx < x; j++ {
			total += beratIkan[idx]
			idx++
		}
		totalPerWadah[i] = total
	}

	fmt.Println("Total berat per wadah:")
	for i := 0; i < totalWadah; i++ {
		fmt.Printf("%.2f ", totalPerWadah[i])
	}
	fmt.Println()

	var totalSemua float64 = 0
	for i := 0; i < totalWadah; i++ {
		totalSemua += totalPerWadah[i]
	}
	rataRata := totalSemua / float64(totalWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}
