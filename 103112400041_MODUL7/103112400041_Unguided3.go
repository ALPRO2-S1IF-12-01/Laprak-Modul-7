//BERTHA ADELA
//103112400041

package main

import "fmt"

type arrBalita [100]float64
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i := 1; arrBerat[i] != -1.0; i++ {
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
	}
}
func rerata(arrBerat arrBalita) float64 {
	total := 0.0
	jumlah := 0
	for i := 0; arrBerat[i] != -1.0; i++ {
		total += arrBerat[i]
		jumlah=jumlah+1
	}
	if jumlah == 0 {
		return 0
	}
	return total / float64(jumlah)
}
func main() {
	var bb arrBalita
	var jumlahBalita int
	var berat float64
	var beratMin, beratMax float64
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&jumlahBalita)

	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scanln(&berat)
		bb[i] = berat
	}
	bb[jumlahBalita] = -1.0
	hitungMinMax(bb, &beratMin, &beratMax)
	rerata := rerata(bb)
	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}