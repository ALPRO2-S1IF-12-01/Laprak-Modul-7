// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
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
	for i := 1; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	fmt.Printf("\nBerat kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", max)
}
