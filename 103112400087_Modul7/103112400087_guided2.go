package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("masukkan jumlah buku dan jumlah buku per rak: ")
	fmt.Scan(&x, &y)

	var hargabuku [500]float64
	fmt.Println("\nmasukkan harga setiap buku (dalam ribuan Rp): ")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargabuku[i])
	}

	var hargaratarata []float64
	for i := 0; i < x; i += y {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargabuku[j]
		}
		hargaratarata = append(hargaratarata, total/float64(y))

	}

	min, max := hargabuku[0], hargabuku[0]
	for _, harga := range hargabuku[:x] {
		if harga < min {
			min = harga
		}
		if harga > max {
			max = harga
		}
	}
	fmt.Printf("\nRata-rata harga per rak: ")
	for _, avg := range hargaratarata {
		fmt.Printf("%.2f", avg)
	}
	fmt.Printf("\nHarga Termahal: %.2f Rp\n ", max)
	fmt.Printf("\nHarga Termurah: %.2f Rp\n ", min)
}
