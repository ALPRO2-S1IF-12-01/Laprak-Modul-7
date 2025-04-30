package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukan jumlah buku dan jumlah buku per rak: ")
	fmt.Scan(&x, &y)
	var hargaBuku [500]float64
	fmt.Println("\nMasukan harga setiap buku (dalam ribuan Rp):")
	for i := 0; i < x; i++ {
		fmt.Scan(&hargaBuku[i])

	}
	var hargaRataRata []float64
	for i := 0; i < x; i++ {
		total := 0.0
		for j := i; j < i+y && j < x; j++ {
			total += hargaBuku[j]

		}
		hargaRataRata = append(hargaRataRata, total/float64(y))
	}
	min, max := hargaBuku[0], hargaBuku[0]
	for i := 1; i < x; i++ {
		if hargaBuku[i] < min {
			min = hargaBuku[i]
		}
		if hargaBuku[i] > max {
			max = hargaBuku[i]
		}
	}

	fmt.Printf("\nRata-rata harga per rak: ")
	for _, avg := range hargaRataRata {
		fmt.Printf("%.2f ", avg)
	}
	fmt.Printf("\nHarga termahal: %.2f\n", max)
	fmt.Printf("Harga termurah: %.2f\n", min)
}
