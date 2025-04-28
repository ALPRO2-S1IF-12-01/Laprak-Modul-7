// Felix Pedrosa V

package main

import "fmt"

// Tipe data untuk menyimpan berat balita
type arrBalita [100]float64

// Fungsi untuk menghitung nilai minimum dan maksimum
func hitungMinMax(arr arrBalita, bMin, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for _, berat := range arr {
		if berat != 0 { // Hanya hitung yang bukan nol
			if berat < *bMin {
				*bMin = berat
			}
			if berat > *bMax {
				*bMax = berat
			}
		}
	}
}

// Fungsi untuk menghitung dan mengembalikan rerata
func rerata(arr arrBalita) float64 {
	var total float64
	var count int

	for _, berat := range arr {
		if berat != 0 { // Hanya hitung yang bukan nol
			total += berat
			count++
		}
	}

	if count == 0 {
		return 0 // Menghindari pembagian dengan nol
	}
	return total / float64(count)
}

func main() {
	var data arrBalita
	var jumlahData int

	// Input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&jumlahData)

	// Input berat balita
	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}

	var min, max float64
	hitungMinMax(data, &min, &max)
	rata := rerata(data)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}