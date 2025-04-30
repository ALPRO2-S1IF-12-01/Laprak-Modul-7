//Anastasia Adinda Narendra Indrianto
//103112400085
package main

import (
	"fmt"
)

func main() {
	var anastasiaN int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&anastasiaN)

	if anastasiaN <= 0 || anastasiaN > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	var berat [1000]float64

	fmt.Println("Masukkan berat anak kelinci satu per satu:")
	for i := 0; i < anastasiaN; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < anastasiaN; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
