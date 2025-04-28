// Felix Pedrosa V

package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000.")
		return
	}

	var beratKelinci [1000]float64
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&beratKelinci[i])
	}

	terkecil := math.MaxFloat64
	terbesar := -math.MaxFloat64

	for i := 0; i < N; i++ {
		if beratKelinci[i] < terkecil {
			terkecil = beratKelinci[i]
		}
		if beratKelinci[i] > terbesar {
			terbesar = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f\n", terkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", terbesar)
}