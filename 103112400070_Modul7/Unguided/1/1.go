// Achmad Zulvan Nur Hakim 103112400070 
package main

import "fmt"


func main() {
	var K int
	fmt.Print("Jumlah kelinci: ")
	fmt.Scan(&K)

	beratK := make([]float64, K)

	for i := 0; i < K; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&beratK[i])
	}

	min, max := beratK[0], beratK[0]
	for i := 1; i < K; i++ {
		if beratK[i] < min {
			min = beratK[i]
		}
		if beratK[i] > max {
			max = beratK[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}