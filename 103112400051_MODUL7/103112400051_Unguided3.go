package main

import "fmt"

func hitungMinMax(arr []float64, n int, bMin *float64, bMax *float64) {
	min := arr[0]
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	*bMin = min
	*bMax = max
}

func rerata(arr []float64, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scanln(&n)

	var berat [100]float64

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	var min, max float64
	hitungMinMax(berat[:], n, &min, &max)
	rata := rerata(berat[:], n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
