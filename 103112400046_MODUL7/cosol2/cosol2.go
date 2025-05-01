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
		fmt.Printf("Harga buku ke-%d: ", i+1)
		fmt.Scan(&hargaBuku[i])
	}

	var hargaRataRata []float64
	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]
		}
		rata := total / float64(min(y, x-i))
		hargaRataRata = append(hargaRataRata, rata)
	}

	fmt.Println("\nRata-rata harga per rak:")
	for i, rata := range hargaRataRata {
		fmt.Printf("Rak ke-%d: %.2f ribu rupiah\n", i+1, rata)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

