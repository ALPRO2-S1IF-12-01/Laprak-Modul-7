//RAJA MUHAMMAD LUFHTI_103112400027
package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arr arrBalita, jumlah int, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < jumlah; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func rerata(arr arrBalita, jumlah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlah; i++ {
		total += arr[i]
	}
	return total / float64(jumlah)
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

	var min, max float64
	hitungMinMax(data, n, &min, &max)
	rerataBerat := rerata(data, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBerat)
}
