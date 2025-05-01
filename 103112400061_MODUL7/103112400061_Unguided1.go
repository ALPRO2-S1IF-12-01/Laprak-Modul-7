package main

import "fmt"

func main() {
	var (
		n int
	)

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	var kandangKelinci [1000]float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d (kg): ",i+1)
		fmt.Scan(&kandangKelinci[i])
	}

	min, max := kandangKelinci[0], kandangKelinci[0]
	for i := 0; i < n; i++ {
		if kandangKelinci[i] < min {
			min = kandangKelinci[i]
		}
		if kandangKelinci[i] > max {
			max = kandangKelinci[i]
		}
	}

	fmt.Printf("\nBerat anak kelinci terberat: %.2f kg\n", max)
	fmt.Printf("Berat anak kelinci teringan: %.2f kg\n", min)
}