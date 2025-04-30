// DWI OKTA SURYANIGNRUM
// 103112400066

package main

import "fmt"

// Mendefinisikan tipe arrBalita sebagai array dengan ukuran maksimum 100 elemen
type arrBalita [100]float64

// Fungsi untuk menghitung berat minimum dan maksimum dari array arrBerat
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	// Menginisialisasi bMin dan bMax dengan nilai berat balita pertama
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Loop untuk mencari nilai minimum dan maksimum
	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i] // Jika ditemukan yang lebih kecil, set bMin
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i] // Jika ditemukan yang lebih besar, set bMax
			}
		}
	}
	
	// Menampilkan hasil minimum dan maksimum
	fmt.Printf("Berat balita minimum: %.2f kg\n", *bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", *bMax)
}

// Fungsi untuk menghitung rata-rata berat balita
func rerata(arrBerat arrBalita) float64 {
	var pembagi, jumlah float64

	// Menghitung total berat dan jumlah data yang valid
	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] > 0 {
			jumlah += arrBerat[i]
			pembagi++
		}
	}

	// Menghitung dan mengembalikan rata-rata
	return jumlah / pembagi
}

func main() {
	var N int
	var beratBalita arrBalita // Array untuk menampung data berat balita

	// Meminta input dari pengguna tentang jumlah balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&N)

	// Memasukkan berat balita ke dalam array
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	// Menghitung berat minimum dan maksimum
	var min, max float64
	hitungMinMax(beratBalita, &min, &max)

	// Menghitung dan menampilkan rerata berat balita
	fmt.Printf("Rata-rata: %.2f\n", rerata(beratBalita))
}