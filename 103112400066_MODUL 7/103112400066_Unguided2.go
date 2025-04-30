// DWI OKTA SURYANINGRUM
// 103112400066

package main

import "fmt"

func main() {
	var x, y int                   // x = jumlah ikan yang akan dihitung, y = kapasitas wadah
	var berat [1000]float64        // Array untuk menyimpan berat tiap ikan, kapasitas maksimum 1000 ikan

	// Meminta pengguna untuk memasukkan jumlah ikan dan kapasitas wadah
	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)               // Membaca jumlah ikan dan kapasitas wadah yang dimasukkan pengguna

	// Meminta pengguna untuk memasukkan berat masing-masing ikan
	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])        // Membaca berat setiap ikan dan menyimpannya ke dalam array
	}

	// Menyiapkan wadah untuk menampung total berat ikan per wadah
	var totalWadah []float64      // Slice untuk menyimpan total berat ikan dalam setiap wadah
	currentTotal := 0.0           // Variabel untuk menghitung total berat ikan dalam wadah saat ini
	count := 0                    // Variabel untuk menghitung jumlah ikan dalam satu wadah

	// Mengelompokkan ikan ke dalam wadah berdasarkan kapasitas dan menghitung total berat per wadah
	for i := 0; i < x; i++ {
		currentTotal += berat[i]  // Menambahkan berat ikan ke total berat wadah saat ini
		count++                   // Menambahkan jumlah ikan yang sudah dimasukkan ke wadah saat ini

		// Jika wadah sudah penuh atau ini adalah ikan terakhir, simpan total wadah dan reset
		if count == y || i == x-1 {               
			totalWadah = append(totalWadah, currentTotal) // Simpan total berat wadah ke dalam slice
			currentTotal = 0.0                  // Reset total berat untuk wadah berikutnya
			count = 0                           // Reset jumlah ikan untuk wadah berikutnya
		}
	}

	// Menampilkan total berat per wadah
	fmt.Println("Total berat per wadah:")
	for _, total := range totalWadah {
		fmt.Printf("%.2f ", total)   // Menampilkan total berat ikan dalam setiap wadah dengan format 2 desimal
	}
	fmt.Println()

	// Menghitung rata-rata berat ikan per wadah
	var sum float64
	for _, total := range totalWadah {
		sum += total                // Menjumlahkan berat total setiap wadah
	}
	average := sum / float64(len(totalWadah)) // Menghitung rata-rata dari semua total wadah

	// Menampilkan hasil rata-rata berat per wadah
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", average)
}
