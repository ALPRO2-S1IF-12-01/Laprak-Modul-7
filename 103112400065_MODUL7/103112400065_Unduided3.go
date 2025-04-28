package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin, *bMax = arrBerat[0], arrBerat[0]
	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
	fmt.Printf("Berat balita minimum: %.2f kg\n", *bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *bMax)
}

func rerata(arrBerat arrBalita) float64 {
	var pembagi, jumlah float64
	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] > 0 {
			jumlah += arrBerat[i]
			pembagi++
		}
	}
	return jumlah / pembagi
}

/*
Dimas Ramadhani
103112400065
*/
func main() {
	var dataBalita arrBalita
	var min, max float64
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}
	hitungMinMax(dataBalita, &min, &max)
	fmt.Printf("Rata-rata: %.2f", rerata(dataBalita))
}
