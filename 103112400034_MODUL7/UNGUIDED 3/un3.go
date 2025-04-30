package main

import "fmt"

func hitungMinMax(arr []float64, bMin *float64, bMax *float64) {
	*bMin = arr[0]
	*bMax = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < *bMin {
			*bMin = arr[i]
		}
		if arr[i] > *bMax {
			*bMax = arr[i]
		}
	}
}

func hitungRerata(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += v
	}
	return total / float64(len(arr))
}

func main() {
	var n int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	var min, max float64
	hitungMinMax(arr, &min, &max)
	rerata := hitungRerata(arr)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
