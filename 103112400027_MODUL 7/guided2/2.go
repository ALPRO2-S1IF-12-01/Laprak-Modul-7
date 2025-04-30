package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukkan jumlah buku dan jumlah buku per rak: ")
	fmt.Scan(&x, &y)

	var hargaBuku [500]float64
	fmt.Println("\nmasukkan setiap harga buku(dalam ribuan):")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargaBuku[i])
	}

	var hargaRataRata []float64
	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]
		}
		hargaRataRata = append(hargaRataRata, total/float64(y))
	}

	min, max := hargaBuku[0], hargaBuku[0]
	for _, harga := range hargaBuku {
		if harga < min {
			min = harga
		}
		if harga > max {
			max = harga
		}
	}

	fmt.Printf("\nRata-rata harga buku per rak:")
	for _, avg := range hargaRataRata {
		fmt.Printf(" %.2f", avg )
	}
	fmt.Printf("\nHarga buku termurah: %.2f", min)	
	fmt.Printf("\nHarga buku termahal: %.2f\n", max)
}