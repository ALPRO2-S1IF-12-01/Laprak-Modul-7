package main
// Setyo Nugroho
// 103112400024

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, jumlah int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < jumlah; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, jumlah int) float64 {
	var total float64 = 0.0
	for i := 0; i < jumlah; i++ {
		total += arrBerat[i]
	}
	return total / float64(jumlah)
}

func main() {
	var dataBerat arrBalita
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah balita harus antara 1 sampai 100.")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBerat[i])
	}

	var beratMin, beratMax float64
	hitungMinMax(dataBerat, n, &beratMin, &beratMax)
	rerataBerat := rerata(dataBerat, n)

	fmt.Printf("\nBerat balita minimum: %.2f kg\n", beratMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", beratMax)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerataBerat)
}