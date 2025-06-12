package main

import "fmt"

func main() {
	fmt.Println("Program Distribusi Ikan ke Wadah")
	fmt.Println("==============================")

	var jumlahIkan, kapasitasWadah int
	var berat [1000]float64

	fmt.Print("\nMasukkan jumlah ikan: ")
	fmt.Scan(&jumlahIkan)

	fmt.Print("Masukkan kapasitas per wadah: ")
	fmt.Scan(&kapasitasWadah)

	if jumlahIkan <= 0 || jumlahIkan > 1000 || kapasitasWadah <= 0 {
		fmt.Println("Error: Input tidak valid!")
		fmt.Println("- Jumlah ikan harus antara 1-1000")
		fmt.Println("- Kapasitas wadah harus lebih dari 0")
		return
	}

	fmt.Println("\nMasukkan berat ikan (dalam kg):")
	for i := 0; i < jumlahIkan; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])

		if berat[i] <= 0 {
			fmt.Println("Error: Berat ikan harus lebih dari 0!")
			return
		}
	}

	jumlahWadah := (jumlahIkan + kapasitasWadah - 1) / kapasitasWadah
	totalBerat := make([]float64, jumlahWadah)

	for i := 0; i < jumlahIkan; i++ {
		wadah := i / kapasitasWadah
		totalBerat[wadah] += berat[i]
	}

	fmt.Println("\nHasil Distribusi Ikan")
	fmt.Println("====================")
	fmt.Printf("Jumlah ikan: %d\n", jumlahIkan)
	fmt.Printf("Kapasitas per wadah: %d\n", kapasitasWadah)
	fmt.Printf("Jumlah wadah: %d\n\n", jumlahWadah)

	fmt.Println("Berat Ikan per Wadah:")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, totalBerat[i])
	}

	var total float64
	for i := 0; i < jumlahWadah; i++ {
		total += totalBerat[i]
	}

	fmt.Printf("\nRata-rata berat per wadah: %.2f kg\n", total/float64(jumlahWadah))
}
