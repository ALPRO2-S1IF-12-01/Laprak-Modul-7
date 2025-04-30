//Anastasia Adinda Narendra Indrianto
//103112400085
package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, jumlahData int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < jumlahData; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, jumlahData int) float64 {
	var total float64
	for i := 0; i < jumlahData; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlahData)
}

func main() {
	var jumlahData int
	var berat arrBalita

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jumlahData)

	if jumlahData > 100 {
		fmt.Println("Jumlah data tidak boleh lebih dari 100.")
		return
	}

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	var bMin, bMax float64
	hitungMinMax(berat, jumlahData, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)

	rata := rerata(berat, jumlahData)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
