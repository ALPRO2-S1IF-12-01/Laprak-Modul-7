package main

import "fmt"

func main() {
	var beratIkan [1000]float64
	var x, y int

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan tidak valid")
		return
	}

	if y <= 0 {
		fmt.Println("Kapasitas wadah tidak valid")
		return
	}

	fmt.Println("Masukkan berat ikan (kg):")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalWadah []float64
	var totalSemua float64
	jumlahWadah := 0

	for i := 0; i < x; i += y {
		total := 0.0
		count := 0
		for j := i; j < i+y && j < x; j++ {
			total += beratIkan[j]
			count++
		}
		totalWadah = append(totalWadah, total)
		totalSemua += total
		jumlahWadah++
	}

	fmt.Println("\nTotal berat per wadah:")
	for _, berat := range totalWadah {
		fmt.Printf("%.2f ", berat)
	}

	if jumlahWadah > 0 {
		rataRata := totalSemua / float64(jumlahWadah)
		fmt.Printf("\nRata-rata berat per wadah: %.2f kg\n", rataRata)
	} else {
		fmt.Println("\nTidak ada wadah yang terisi")
	}
}
