package main

import (
	"fmt"
)

func main() {
	var berat [1000]float64
	var x, y int

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalPerWadah := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		wadahKe := i / y
		totalPerWadah[wadahKe] += berat[i]
	}

	total := 0.0
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah ke-%d: %.2f kg\n", i+1, totalPerWadah[i])
		total += totalPerWadah[i]
	}

	rata2 := total / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f kg\n", rata2)
}
