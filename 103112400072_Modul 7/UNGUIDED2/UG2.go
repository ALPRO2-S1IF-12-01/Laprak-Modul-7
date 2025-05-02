package main

import (
	"fmt"
)

func main() {

	var beratIkan [1000]float64

	var x, y int
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan harus antara 1-1000")
		return
	}

	if y <= 0 {
		fmt.Println("Jumlah wadah harus positif")
		return
	}

	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalBeratPerWadah [1000]float64
	var ikanPerWadah [1000]int

	for i := 0; i < x; i++ {

		wadahIdx := i % y
		totalBeratPerWadah[wadahIdx] += beratIkan[i]
		ikanPerWadah[wadahIdx]++
	}

	for i := 0; i < y; i++ {
		fmt.Printf("%.2f", totalBeratPerWadah[i])
		if i < y-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	var totalRataRata float64
	for i := 0; i < y; i++ {
		var rataRata float64
		if ikanPerWadah[i] > 0 {
			rataRata = totalBeratPerWadah[i] / float64(ikanPerWadah[i])
		}
		totalRataRata += rataRata
	}

	rataRataKeseluruhan := totalRataRata / float64(y)
	fmt.Printf("%.2f\n", rataRataKeseluruhan)
}
