//RAJA MUHAMMAD LUFHTI_103112400027
package main

import (
	"fmt"
)

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}

	var totalWadah [1000]float64
	jumlahWadah := (x + y - 1) / y 

	for i := 0; i < x; i++ {
		indexWadah := i / y
		totalWadah[indexWadah] += ikan[i]
	}

	fmt.Print("Total berat di setiap wadah:\n")
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", totalWadah[i])
	}
	fmt.Println()

	var total float64
	for i := 0; i < jumlahWadah; i++ {
		total += totalWadah[i]
	}
	rataRata := total / float64(jumlahWadah)
	fmt.Printf("Berat rata-rata per wadah: %.2f\n", rataRata)
}
