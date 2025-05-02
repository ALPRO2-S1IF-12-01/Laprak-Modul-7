package main

import (
	"fmt"
)

func main() {

	var beratKelinci [1000]float64

	var n int
	fmt.Scan(&n)

	if n <= 0 || n > 1000 {
		fmt.Println("Jumlah kelinci harus antara 1-1000")
		return
	}

	for i := 0; i < n; i++ {
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

	fmt.Printf("%.2f %.2f\n", beratTerkecil, beratTerbesar)
}
