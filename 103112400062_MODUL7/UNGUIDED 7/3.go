//M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
)

type arrBalita [100]float64

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

func rataRata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("masukkan jumlah balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)

	rataRata := rataRata(berat, n)

	fmt.Printf("berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("rata-rata berat balita: %.2f kg\n", rataRata)
}