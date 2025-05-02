package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {

	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := range arrBerat {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {

	total := 0.0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}

	if n > 0 {
		return total / float64(n)
	}
	return 0
}

func main() {
	var dataBalita arrBalita
	var n int

	fmt.Print("Masukkan banyak data berat balita : ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1-100")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	var min, max float64
	hitungMinMax(dataBalita, &min, &max)

	rataRata := rerata(dataBalita, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
