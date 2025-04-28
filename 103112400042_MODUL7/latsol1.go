// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah anak kelinci yang akan ditimbang: ")
	fmt.Scan(&n)

	var beratKelinci [1000]float64
	fmt.Println("\nMasukkan berat setiap anak kelinci (dalam kg):")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	min, max := beratKelinci[0], beratKelinci[0]
	total := 0.0
	for i := 0; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
		total += beratKelinci[i]
	}

	rataRata := total / float64(n)

	fmt.Printf("\nBerat kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat kelinci: %.2f kg\n", rataRata)
}
