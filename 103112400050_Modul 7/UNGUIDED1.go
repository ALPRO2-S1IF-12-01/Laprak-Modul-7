package main

import (
	"fmt"
)

// 103112400050
func main() {
	var beratKelinci [1000]float64
	var n int

	fmt.Print("Masukkan jumlah anak kelinci (maksimal 1000): ")
	fmt.Scan(&n)

	if n < 1 || n > 1000 {
		fmt.Println("Jumlah tidak valid! Harus antara 1 hingga 1000.")
		return
	}

	fmt.Println("\nMasukkan berat masing-masing anak kelinci:")
	for i := 0; i < n; i++ {
		fmt.Printf("Berat anak kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]

	for i := 1; i < n; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("\nBerat terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f kg\n", beratTerbesar)
}
