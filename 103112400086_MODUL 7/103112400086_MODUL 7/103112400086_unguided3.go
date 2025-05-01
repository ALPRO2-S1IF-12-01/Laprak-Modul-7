package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah balita: ")
	fmt.Scan(&n)

	var berat []float64
	var total, min, max float64

	for i := 0; i < n; i++ {
		var b float64
		fmt.Printf("Masukkan berat balita ke-%d (kg): ", i+1)
		fmt.Scan(&b)
		berat = append(berat, b)

		if i == 0 {
			min = b
			max = b
		} else {
			if b < min {
				min = b
			}
			if b > max {
				max = b
			}
		}
		total += b
	}

	rata := total / float64(n)

	fmt.Printf("\nBerat balita terkecil: %.2f kg\n", min)
	fmt.Printf("Berat balita terbesar: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
