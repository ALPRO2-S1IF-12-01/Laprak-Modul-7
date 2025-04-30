package main
// Setyo Nugroho
// 103112400024

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64
	fmt.Print("Masukkan jumlah ikan dan jumlah ikan per wadah: ")
	fmt.Scan(&x, &y)
	fmt.Println("Masukkan berat tiap ikan: ")

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratWadah := make([]float64, jumlahWadah)
	ikanDiWadah := make([]int, jumlahWadah)

	for i := 0; i < x; i++ {
		wadah := i / y
		totalBeratWadah[wadah] += ikan[i]
		ikanDiWadah[wadah]++
	}

	fmt.Println("\nTotal Berat per Wadah")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah Ke %d - Total Berat: %.2f kg\n", i+1, totalBeratWadah[i])
	}

	fmt.Println("\nRata rata Berat per Wadah")
	for i := 0; i < jumlahWadah; i++ {
		rataRata := totalBeratWadah[i] / float64(ikanDiWadah[i])
		fmt.Printf("Wadah Ke %d Rata rata Berat: %.2f kg\n", i+1, rataRata)
	}
}