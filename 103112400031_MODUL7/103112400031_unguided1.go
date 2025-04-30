// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	var beratKelinci [1000]float64
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%v (kg): ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	min, max := beratKelinci[0], beratKelinci[0]
	for i := 1; i < N; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}
	fmt.Printf("\nBerat anak kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat anak kelinci terberat: %.2f kg\n", max)
}