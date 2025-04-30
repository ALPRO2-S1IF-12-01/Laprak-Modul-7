// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func main() {
	var x, y int
	var beratIkan [1000]float64

	fmt.Print("Jumlah ikan : ")
	fmt.Scan(&x)
	fmt.Print("Kapasitas wadah : ")
	fmt.Scan(&y)
	
	for i := 0; i < x; i++ {
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	jmlWadah := x / y
	if x%y != 0 {
		jmlWadah++
	}

	totalPWadah := make([]float64, jmlWadah)
	for i := 0; i < x; i++ {
		wadah := i / y
		totalPWadah[wadah] += beratIkan[i]
	}

	fmt.Print("Total berat per wadah: ")
	for i := 0; i < jmlWadah; i++ {
		fmt.Printf("%.2f ", totalPWadah[i])
	}
	fmt.Println()

	var total float64
	for i := 0; i < x; i++ {
		total += beratIkan[i]
	}
	avg := total / float64(x)

	fmt.Printf("Rata-rata berat ikan: %.2f\n ", avg)
}	