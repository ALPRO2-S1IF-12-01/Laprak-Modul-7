package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64
	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	fmt.Print("Masukkan berat setiap ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	banyakWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, banyakWadah)
	jumlahIkanWadah := make([]float64, banyakWadah)
	for dimasRamadhani := 0; dimasRamadhani < x; dimasRamadhani++ {
		posisiWadah := dimasRamadhani / y
		totalBeratWadah[posisiWadah] += ikan[dimasRamadhani]
		jumlahIkanWadah[posisiWadah]++
	}

	for i, beratWadah := range totalBeratWadah {
		fmt.Printf("| Wadah ke-%d: %.2f |", i+1, beratWadah)
	}
	// 103112400065
	fmt.Println()

	for i := 0; i < banyakWadah; i++ {
		rataRata := totalBeratWadah[i] / jumlahIkanWadah[i]
		fmt.Printf("| Berat rata-rata wadah ke-%d: %.2f |", i+1, rataRata)
	}
}
