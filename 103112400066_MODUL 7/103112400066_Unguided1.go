// DWI OKTA SURYANINGRUM
// 103112400066

package main

import (
	"fmt"
)

func main() {
	var N int                      // Menyimpan jumlah anak kelinci yang akan ditimbang
	var berat [1000]float64       // Array untuk menyimpan berat anak kelinci (maksimal 1000 data)

	// Input jumlah anak kelinci dari pengguna
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)                  // Membaca jumlah anak kelinci

	// Input berat masing-masing anak kelinci sebanyak N kali
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])       // Membaca berat anak kelinci dan menyimpannya ke array
	}

	// Inisialisasi nilai terkecil dan terbesar dengan elemen pertama
	terkecil := berat[0]          // Asumsikan berat pertama sebagai yang terkecil
	terbesar := berat[0]          // Asumsikan berat pertama sebagai yang terbesar

	// Loop untuk mencari nilai terkecil dan terbesar
	for i := 1; i < N; i++ {
		if berat[i] < terkecil {  // Jika berat ke-i lebih kecil dari yang terkecil sebelumnya
			terkecil = berat[i]   // Perbarui nilai terkecil
		}
		if berat[i] > terbesar {  // Jika berat ke-i lebih besar dari yang terbesar sebelumnya
			terbesar = berat[i]   // Perbarui nilai terbesar
		}
	}

	// Menampilkan hasil akhir: berat terkecil dan terbesar
	fmt.Printf("Berat terkecil: %.2f\n", terkecil)   // Output berat terkecil dengan 2 desimal
	fmt.Printf("Berat terbesar: %.2f\n", terbesar)   // Output berat terbesar dengan 2 desimal
}