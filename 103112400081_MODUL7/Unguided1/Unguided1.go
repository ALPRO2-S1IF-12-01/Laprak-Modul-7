// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	weights := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&weights[i])
	}
	sort.Float64s(weights)

	fmt.Printf("Berat terkecil: %.2f\n", weights[0])
	fmt.Printf("Berat terbesar: %.2f\n", weights[N-1])
}
