package main

import "fmt"

func hitungMinMax(arr []float64, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func hitungRerata(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}

func main() {
	fmt.Println("Program Analisis Berat Badan Balita")
	fmt.Println("=================================")

	var jumlahBalita int
	fmt.Print("\nMasukkan jumlah balita: ")
	fmt.Scan(&jumlahBalita)

	if jumlahBalita <= 0 {
		fmt.Println("Error: Jumlah balita harus lebih dari 0!")
		return
	}

	beratBalita := make([]float64, jumlahBalita)

	fmt.Println("\nMasukkan berat badan balita (dalam kg):")
	for i := 0; i < jumlahBalita; i++ {
		fmt.Printf("Berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])

		if beratBalita[i] <= 0 {
			fmt.Println("Error: Berat badan harus lebih dari 0!")
			return
		}
	}

	var min, max float64
	hitungMinMax(beratBalita, &min, &max)
	rerata := hitungRerata(beratBalita)

	fmt.Println("\nHasil Analisis Berat Badan Balita")
	fmt.Println("================================")
	fmt.Printf("Jumlah balita: %d\n", jumlahBalita)
	fmt.Printf("Berat minimum: %.2f kg\n", min)
	fmt.Printf("Berat maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat: %.2f kg\n", rerata)
	var diBawahRerata, diAtasRerata int
	for _, berat := range beratBalita {
		if berat < rerata {
			diBawahRerata++
		} else if berat > rerata {
			diAtasRerata++
		}
	}

	fmt.Printf("\nBalita dengan berat di bawah rata-rata: %d\n", diBawahRerata)
	fmt.Printf("Balita dengan berat di atas rata-rata: %d\n", diAtasRerata)
}
