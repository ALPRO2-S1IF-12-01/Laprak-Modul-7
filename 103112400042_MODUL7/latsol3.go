// M.HANIF AL FAIZ
// 103112400042
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

func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func countAboveAverage(arrBerat arrBalita, n int, rataRata float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if arrBerat[i] > rataRata {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &bMin, &bMax)

	rataRata := rerata(berat, n)

	count := countAboveAverage(berat, n, rataRata)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
	fmt.Printf("Jumlah balita dengan berat di atas rata-rata: %d\n", count)
}
