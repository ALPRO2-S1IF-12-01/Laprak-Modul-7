//ABID FADHILAH M
//103112400046
package main

import (
	"fmt"
)

func main() {
	var n int
	var berat [100]float64
	var bMin, bMax, total float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
		if i == 0 || berat[i] < bMin {
			bMin = berat[i]
		}
		if berat[i] > bMax {
			bMax = berat[i]
		}
		total += berat[i]
	}

	rataRata := total / float64(n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
