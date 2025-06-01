//Nama : Anggun Wahyu Widiyana (103112480280)
package main

import (
	"fmt"
)

type dataBalita [100]float64

func cariMinMax(data dataBalita, jumlah int, minBerat, maxBerat *float64) {
	*minBerat = data[0]
	*maxBerat = data[0]
	for i := 1; i < jumlah; i++ {
		if data[i] < *minBerat {
			*minBerat = data[i]
		}
		if data[i] > *maxBerat {
			*maxBerat = data[i]
		}
	}
}

func hitungRerata(data dataBalita, jumlah int) float64 {
	total := 0.0
	for i := 0; i < jumlah; i++ {
		total += data[i]
	}
	return total / float64(jumlah)
}

func main() {
	var n int
	var berat dataBalita

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	var min, max float64
	cariMinMax(berat, n, &min, &max)
	rata := hitungRerata(berat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}