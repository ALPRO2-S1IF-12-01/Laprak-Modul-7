//ABID FDAHILAH M
//103112400046
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scan(&n)

	var berat [1000]float64
	min, max := 0.0, 0.0

	for i := 0; i < n; i++ {
		fmt.Printf("Berat ke-%d (kg): ", i+1)
		fmt.Scan(&berat[i])
		if i == 0 || berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Printf("\nTerkecil: %.2f kg\nTerbesar: %.2f kg\n", min, max)
}
