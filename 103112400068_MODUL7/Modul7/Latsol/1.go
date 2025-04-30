package main

import "fmt"

func main() {

	var beratKelinci [1000]float64
	var n int

	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah kelinci tidak valid")
		return
	}

	fmt.Println("Masukkan berat kelinci (kg):")
	for i := 0; i < n; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	min, max := beratKelinci[0], beratKelinci[0]

	for i := 1; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", max)
}
