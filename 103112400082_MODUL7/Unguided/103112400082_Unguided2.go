//RizkinaAzizah 103112400082
package main

import "fmt"

func main() {
	var jumlahIkan, jumlahWadah int
	const maxCapacity = 1000
	var beratIkan [maxCapacity]float64

	fmt.Scan(&jumlahIkan, &jumlahWadah)

	for i := 0; i < jumlahIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	var totalBerat [maxCapacity]float64
	var jumlahDalamWadah [maxCapacity]int

	for i := 0; i < jumlahIkan; i++ {
		indeksWadah := i / jumlahWadah
		totalBerat[indeksWadah] += beratIkan[i]
		jumlahDalamWadah[indeksWadah]++
	}

	for i := 0; i < (jumlahIkan+jumlahWadah-1)/jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalBerat[i])
	}
	fmt.Println()

	for i := 0; i < (jumlahIkan+jumlahWadah-1)/jumlahWadah; i++ {
		if jumlahDalamWadah[i] > 0 {
			beratRata := totalBerat[i] / float64(jumlahDalamWadah[i])
			fmt.Printf("%.2f ", beratRata)
		} else {
			fmt.Printf("0.00 ")
		}
	}
	fmt.Println()
}
