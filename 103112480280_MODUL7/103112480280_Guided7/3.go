// Nama : Anggun Wahyu Widiyana (103112480280)
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan jumlah siswa: ")
	fmt.Scan(&n)

	var nilaiSiswa [200]float64
	var totalNilai float64 = 0

	fmt.Println("\nMasukan nilai ujian masing-masing siswa:")
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

	fmt.Printf("\nNilai terendah: %.0f\n", min)
	fmt.Printf("Nilai terendah: %.0f\n", max)
	fmt.Printf("Rata-rata kelas: %.2f\n", rataRata)
	fmt.Printf("Jumlah siswa di atas rata-rata: %d\n", diAtasRataRata)
}
