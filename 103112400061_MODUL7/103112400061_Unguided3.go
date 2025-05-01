package main

import "fmt"

type arrBalita [100]float64

func main() {
	var (
		n int
		arrBalita arrBalita
	)

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBalita[i])
	}

	beratMin := arrBalita[0]
	beratMax := arrBalita[0]
	
	hitungMinMax(arrBalita, &beratMin, &beratMax)
	rataRata := rerata(arrBalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maximum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) { // 103112400061
	for i := 0; i < len(arrBerat) && arrBerat[i] != 0; i++ {
		if arrBerat[i] != 0 {
			if arrBerat[i] < *bMin {
				*bMin = arrBerat[i]
			}
			if arrBerat[i] > *bMax {
				*bMax = arrBerat[i]
			}
		}
	}
}

func rerata(arrBerat arrBalita) float64 {
	var avg, totalBerat float64
	var n int
	for i := 0; arrBerat[i] != 0; i++ {
		n++
	}
	for _, berat := range arrBerat {
		totalBerat += berat
	}
	avg = totalBerat / float64(n)
	return avg
}