// DWI OKTA SURYANINGRUM
// 1013112400066

package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukan jumlah buku dan jumlah buku per rak:  ")
	fmt.Scan(&x, &y)


	var hargaBuku [500]float64
	fmt.Println("masukan harga setiap buku (dalam ribuan Rp.):")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargaBuku[i])
	}

	var hargaRataRata []float64
	for i := 0; i < x; i+= y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]
		}
		hargaRataRata = append(hargaRataRata, total/float64(y))
	}

	min, max := hargaRataRata[0], hargaRataRata[0]
	for _, harga := range hargaBuku[:x] {
		if harga < min {
			min = harga
		}
		if harga > max {
			max = harga
		}
	}

	fmt.Printf("\n harga rata-rata per rak: ")
	for _, avg := range hargaRataRata {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Printf("\n harga buku termahal: %.2f", max)
	fmt.Printf("\n harga buku termurah: %.2f\n", min)
}