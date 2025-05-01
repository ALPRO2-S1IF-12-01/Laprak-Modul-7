package main

import "fmt"

func main() {
	var (
		x, y int
		totalWadah int
		totalBeratIwak, avg float64
	)
	fmt.Print("Masukkan jumlah ikan yang akan dijual dan ikan yang dimasukkan ke dalam wadah: ")
	fmt.Scan(&x, &y)

	var iwak [1000]float64
	fmt.Println("\nMasukkan berat setiap ikan (kg):")
	for i := 0; i < x; i++ { // 103112400061
		fmt.Scan(&iwak[i])
	}

	totalWadah = (x+y-1) / y
	totalBeratTiapWadah := make([]float64, totalWadah)

	for nope := 0; nope < x; nope++ {
		iWadah := nope / y
		totalBeratTiapWadah[iWadah] += iwak[nope]
	}

	fmt.Print("\nTotal berat ikan tiap wadah: ")
	for _, wadah := range totalBeratTiapWadah {
		totalBeratIwak += wadah
		fmt.Printf("%.2f ", wadah)
	}
	if len(totalBeratTiapWadah) < y {
		fmt.Print(y-y)
	}
	avg = totalBeratIwak / float64(totalWadah)
	fmt.Printf("\nRata-rata berat ikan tiap wadah: %.2f\n", avg)
}