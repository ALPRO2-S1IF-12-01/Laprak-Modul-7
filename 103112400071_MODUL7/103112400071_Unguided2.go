// Raihan Adi Arba
// 103112400071

package main

import "fmt"

func main() {
	var jumlahIkan, jumlahPerWadah int

	// Input jumlah ikan dan jumlah ikan per wadah
	fmt.Print("Masukkan jumlah ikan dan jumlah per wadah: ")
	fmt.Scan(&jumlahIkan, &jumlahPerWadah)

	var ikan [1000]float64

	// Input berat tiap ikan
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var wadah [1000]float64
	jumlahWadah := 0
	totalSemua := 0.0

	// Proses pembagian ke wadah
	for i := 0; i < jumlahIkan; i += jumlahPerWadah {
		batas := i + jumlahPerWadah
		if batas > jumlahIkan {
			batas = jumlahIkan
		}

		total := 0.0
		for j := i; j < batas; j++ {
			total += ikan[j]
		}

		wadah[jumlahWadah] = total
		jumlahWadah++
		totalSemua += total
	}

	// Output berat per wadah
	fmt.Println("\nBerat per wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, wadah[i])
	}

	// Output rata-rata
	rataRata := totalSemua / float64(jumlahIkan)
	fmt.Printf("\nRata-rata berat ikan: %.2f kg\n", rataRata)
}
