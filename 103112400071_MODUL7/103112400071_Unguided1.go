// Raihan Adi Arba
// 103112400071

package main

import "fmt"

func inputBeratKelinci(n int, berat *[1000]float64) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat[i])
	}
}

func cariMinMax(n int, berat [1000]float64) (float64, float64) {
	min := berat[0]
	max := berat[0]

	for i := 1; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	return min, max
}

func main() {
	var jumlahKelinci int
	var beratKelinci [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	inputBeratKelinci(jumlahKelinci, &beratKelinci)

	min, max := cariMinMax(jumlahKelinci, beratKelinci)

	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}
