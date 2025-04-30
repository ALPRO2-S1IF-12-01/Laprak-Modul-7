//Ahmad Ruba'i
//103112400074 
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	
	for i := 1; i < len(arrBerat); i++ {
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
	var total float64 = 0
	var count int = 0
	
	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] != 0 {
			total += arrBerat[i]
			count++
		}
	}
	
	if count > 0 {
		return total / float64(count)
	}
	return 0
}

func main() {
	var n int
	var dataBalita arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}
	hitungMinMax(dataBalita, &min, &max)
	rata := rerata(dataBalita)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}