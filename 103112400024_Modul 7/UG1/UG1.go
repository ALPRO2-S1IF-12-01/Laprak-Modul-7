package main
// Setyo Nugroho
// 103112400024

import "fmt"

func main() {
	var jumlah int
	var dataBerat [1000]float64

	fmt.Print("Input jumlah anak kelinci: ")
	fmt.Scan(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	terberat := dataBerat[0]
	teringan := dataBerat[0]

	for i := 1; i < jumlah; i++ {
		if dataBerat[i] > terberat {
			terberat = dataBerat[i]
		}
		if dataBerat[i] < teringan {
			teringan = dataBerat[i]
		}
	}

	fmt.Printf("Berat anak kelinci paling ringan: %.2f kg\n", teringan)
	fmt.Printf("Berat anak kelinci paling berat : %.2f kg\n", terberat)
}