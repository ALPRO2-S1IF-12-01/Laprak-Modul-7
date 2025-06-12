// daffa tsaqifna f
package main

import "fmt"

type KelinciData struct {
	weights []float64
	count   int
}

const maxSize = 1000

func inputBeratKelinci() KelinciData {
	var n int

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&n)

	if n <= 0 || n > maxSize {
		fmt.Println("Jumlah anak kelinci tidak valid (1 - 1000).")
		return KelinciData{}
	}

	berat := make([]float64, n)
	fmt.Printf("Masukkan %d berat anak kelinci:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	return KelinciData{
		weights: berat,
		count:   n,
	}
}

func cariMinMax(data KelinciData) (float64, float64) {
	if data.count == 0 {
		return 0, 0
	}

	min := data.weights[0]
	max := data.weights[0]

	for i := 1; i < data.count; i++ {
		if data.weights[i] < min {
			min = data.weights[i]
		}
		if data.weights[i] > max {
			max = data.weights[i]
		}
	}

	return min, max
}

func main() {
	data := inputBeratKelinci()

	if data.count == 0 {
		return
	}

	min, max := cariMinMax(data)

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}
