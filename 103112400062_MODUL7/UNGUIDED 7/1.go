//M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukkan jumlah anak kelinci: ")
	fmt.Scanln(&n)

	var beratKelinci [1000]float64
	var beratTerkecil, beratTerbesar float64

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scanln(&beratKelinci[i])

		if i == 0 { 
			beratTerkecil = beratKelinci[i]
			beratTerbesar = beratKelinci[i]
		} else {
			if beratKelinci[i] < beratTerkecil {
				beratTerkecil = beratKelinci[i]
			}
			if beratKelinci[i] > beratTerbesar {
				beratTerbesar = beratKelinci[i]
			}
		}
	}

	fmt.Printf("berat kelinci terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("berat kelinci terbesar: %.2f kg\n", beratTerbesar)
}