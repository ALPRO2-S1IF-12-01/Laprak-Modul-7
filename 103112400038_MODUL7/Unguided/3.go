//Nama	: Hakan Ismail Afnan 
//NIM	: 103112400038
package main

import (
	"fmt"
)

type beratBalita [100]float64

func temukanMinMax(berat beratBalita, banyak int, beratMin, beratMax *float64) {
	*beratMin = berat[0]
	*beratMax = berat[0]
	for i := 1; i < banyak; i++ {
		if berat[i] < *beratMin {
			*beratMin = berat[i]
		}
		if berat[i] > *beratMax {
			*beratMax = berat[i]
		}
	}
}

func rataRataBerat(berat beratBalita, banyak int) float64 {
	jumlah := 0.0
	for i := 0; i < banyak; i++ {
		jumlah += berat[i]
	}
	return jumlah / float64(banyak)
}

func main() {
	var jumlahData int
	var dataBerat beratBalita

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	var minBerat, maxBerat float64
	temukanMinMax(dataBerat, jumlahData, &minBerat, &maxBerat)
	rata := rataRataBerat(dataBerat, jumlahData)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxBerat)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}