// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scanln(&n)

	berat := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scanln(&berat[i])
	}

	minWeight := berat[0]
	maxWeight := berat[0]
	sum := 0.0

	for _, w := range berat {
		minWeight = math.Min(minWeight, w)
		maxWeight = math.Max(maxWeight, w)
		sum += w
	}

	avg := sum / float64(n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minWeight)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maxWeight)
	fmt.Printf("Rerata Berat balita : %.2f kg\n", avg)
}
