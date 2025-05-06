//Muhammad Zaky Mubarok
package main

import (
	"fmt"
)

func hitungMinMax(beratBalita []float64, min *float64, max *float64) {
	if len(beratBalita) == 0 {
		return
	}
	*min = beratBalita[0]
	*max = beratBalita[0]
	for _, berat := range beratBalita[1:] {
		if berat < *min {
			*min = berat
		}
		if berat > *max {
			*max = berat
		}
	}
}

func hitungRerata(beratBalita []float64) float64 {
	if len(beratBalita) == 0 {
		return 0
	}
	var total float64
	for _, berat := range beratBalita {
		total += berat
	}
	return total / float64(len(beratBalita))
}

func main() {
	var jumlah int
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&jumlah)

	beratBalita := make([]float64, jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&beratBalita[i])
	}

	var beratMinimum, beratMaksimum float64
	hitungMinMax(beratBalita, &beratMinimum, &beratMaksimum)
	rerata := hitungRerata(beratBalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMinimum)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMaksimum)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata)
}
