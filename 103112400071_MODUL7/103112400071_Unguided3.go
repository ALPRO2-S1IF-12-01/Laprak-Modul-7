// Raihan Adi Arba
// 103112400071

package main

import (
	"fmt"
)

type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum
func hitungMinMax(arr arrBalita, jumlah int, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < jumlah; i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita
func hitungRerata(arr arrBalita, jumlah int) float64 {
	var total float64
	for i := 0; i < jumlah; i++ {
		total += arr[i]
	}
	return total / float64(jumlah)
}

func main() {
	var data arrBalita
	var n int

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, n, &min, &max)
	rata := hitungRerata(data, n)

	fmt.Println("\nHasil Pengolahan Data Berat Balita:")
	fmt.Printf("• Berat minimum : %.2f kg\n", min)
	fmt.Printf("• Berat maksimum: %.2f kg\n", max)
	fmt.Printf("• Rata-rata     : %.2f kg\n", rata)
}
