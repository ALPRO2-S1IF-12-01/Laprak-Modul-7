package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan jumlah siswa: ")
	fmt.Scan(&n)

	var nilaiSiswa [200]float64
	var totalNilai float64 = 0

	fmt.Println("\nmasukkan nilai ujian masing masing siswa:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nilaiSiswa[i])
		totalNilai += nilaiSiswa[i]
	}

	rataRata := totalNilai / float64(n)

	min, max := nilaiSiswa[0], nilaiSiswa[0]
	var diAtasRataRata int = 0
	for _, nilai := range nilaiSiswa[:n] {
		if nilai < min {
			min = nilai
		}
		if nilai > max {
			max = nilai
		}
		if nilai > rataRata {
			diAtasRataRata++
		}
	}

	fmt.Printf("\nRata-rata nilai kelas: %.2f", rataRata)
	fmt.Printf("\nNilai ujian terendah: %.2f", min)
	fmt.Printf("\nNilai ujian tertinggi: %.2f", max)
	fmt.Printf("\nJumlah siswa di atas rata-rata: %d\n", diAtasRataRata)

}