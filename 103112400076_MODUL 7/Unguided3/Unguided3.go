package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, jumlah int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < jumlah; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// Fungsi untuk menghitung rerata
func rerata(arrBerat arrBalita, jumlah int) float64 {
	var total float64 = 0.0
	for i := 0; i < jumlah; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlah)
}

func main() {
	var dataBerat arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 sampai 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	var beratMin, beratMax float64
	hitungMinMax(dataBerat, n, &beratMin, &beratMax)
	rerataBerat := rerata(dataBerat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBerat)
}
