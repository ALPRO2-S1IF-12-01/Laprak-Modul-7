package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah buku dan jumlah buku per rak: ")
	fmt.Scan(&x, &y)

	var hargaBuku [500]float64
	fmt.Println("\nMasukkan harga setiap buku (dalam ribuan Rp):")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargaBuku[i])
	}

	var hargaRatarata []float64
	for i := 0; i < x; i++ {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]
		}
		hargaRatarata = append(hargaRatarata, total/float64(y))
	}

	min, max := hargaBuku[0], hargaBuku[0]
	for _, harga := range hargaBuku[:x] {
		if harga < min {
			min = harga
		}
		if harga > max {
			max = harga
		}
	}
	fmt.Printf("\nRata-rata harga per rak: ")
	for _, avg := range hargaRatarata {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Printf("\nHarga termahal: Rp %.2f\n", max)
	fmt.Printf("Harga termurah: Rp %.2f\n", min)

}
