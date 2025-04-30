package main

//103112400049 Hisyam Nurdiatmoko

import "fmt"

func main() {
	var x, y int
	var beratIkanHisyam [1000]float64
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah (x y): ")
	fmt.Scan(&x, &y)
	fmt.Print("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkanHisyam[i])
	}
	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}
	totalWadah := make([]float64, jumlahWadah)
	for i := 0; i < x; i++ {
		idx := i / y
		totalWadah[idx] += beratIkanHisyam[i]
	}
	fmt.Println("Total berat tiap wadah:")
	for _, total := range totalWadah {
		fmt.Printf("%.2f ", total)
	}
	fmt.Println()
	var totalSemua float64
	for _, total := range totalWadah {
		totalSemua += total
	}
	rataRata := totalSemua / float64(jumlahWadah)
	fmt.Printf("Rata-rata berat per wadah: %.2f\n", rataRata)
}
