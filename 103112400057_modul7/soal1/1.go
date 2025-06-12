package main

import "fmt"

func main() {
	fmt.Println("Program Pencatatan Berat Kelinci")
	fmt.Println("==============================")

	var jumlahKelinci int
	fmt.Print("\nMasukkan jumlah kelinci: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci <= 0 {
		fmt.Println("Error: Jumlah kelinci harus lebih dari 0!")
		return
	}

	daftarBerat := make([]float64, jumlahKelinci)
	var totalBerat float64

	fmt.Println("\nMasukkan berat kelinci (dalam kg):")
	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&daftarBerat[i])

		if daftarBerat[i] <= 0 {
			fmt.Println("Error: Berat kelinci harus lebih dari 0!")
			return
		}
		totalBerat += daftarBerat[i]
	}

	palingRingan := daftarBerat[0]
	palingBerat := daftarBerat[0]
	indexRingan := 0
	indexBerat := 0

	for i := 1; i < jumlahKelinci; i++ {
		if daftarBerat[i] < palingRingan {
			palingRingan = daftarBerat[i]
			indexRingan = i
		}
		if daftarBerat[i] > palingBerat {
			palingBerat = daftarBerat[i]
			indexBerat = i
		}
	}

	fmt.Println("\nHasil Analisis Berat Kelinci")
	fmt.Println("===========================")
	fmt.Printf("Jumlah kelinci: %d\n", jumlahKelinci)
	fmt.Printf("Rata-rata berat: %.2f kg\n", totalBerat/float64(jumlahKelinci))
	fmt.Printf("Kelinci paling ringan: Kelinci ke-%d (%.2f kg)\n", indexRingan+1, palingRingan)
	fmt.Printf("Kelinci paling berat: Kelinci ke-%d (%.2f kg)\n", indexBerat+1, palingBerat)
}
