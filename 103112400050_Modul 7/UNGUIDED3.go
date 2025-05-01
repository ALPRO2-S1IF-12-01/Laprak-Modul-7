package main

import "fmt"

//103112400050
type arrBalita [100]float64

func hitungMinMax(data arrBalita, jumlah int, min, max *float64) {
	*min = data[0]
	*max = data[0]
	for i := 1; i < jumlah; i++ {
		if data[i] < *min {
			*min = data[i]
		}
		if data[i] > *max {
			*max = data[i]
		}
	}
}

func Rerata(data arrBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += data[i]
	}
	return total / float64(jumlah)
}

func main() {
	var data arrBalita
	var jumlahData int

	fmt.Print("Masukkan jumlah data berat balita: ")
	fmt.Scan(&jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
		fmt.Scan(&data[i])
	}

	var beratMin, beratMax float64
	hitungMinMax(data, jumlahData, &beratMin, &beratMax)
	rataRata := Rerata(data, jumlahData)

	fmt.Printf("\nBerat minimum balita: %.2f kg\n", beratMin)
	fmt.Printf("Berat maksimum balita: %.2f kg\n", beratMax)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rataRata)
}
