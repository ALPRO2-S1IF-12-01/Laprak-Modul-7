//Rizkina Azizah
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, jumlahData int) {
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
	total := 0.0
	for i := 0; i < jumlahData; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlahData)
}

func main() {
	var beratBalita arrBalita
	var jumlahData int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	var beratMin, beratMax float64

	hitungMinMax(beratBalita, &beratMin, &beratMax, jumlahData)

	rata := rerata(beratBalita, jumlahData)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
