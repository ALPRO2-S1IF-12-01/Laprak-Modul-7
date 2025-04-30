package main

import (
	"fmt"
)

const maxData = 100

type arrBalita [maxData]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var data arrBalita
	var n int

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var minBerat, maxBerat float64
	hitungMinMax(data, n, &minBerat, &maxBerat)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", minBerat)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxBerat)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, n))
}
