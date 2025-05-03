package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Jumlah Kelinci: ")
	fmt.Scan(&n)

	var beratKenlinci [100]float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d (Kg): ", i+1)
		fmt.Scan(&beratKenlinci[i])
	}

	min, max := beratKenlinci[0], beratKenlinci[0]
	for i := 1; i < n; i++ {
		if beratKenlinci[i] < min {
			min = beratKenlinci[i]
		}
		if beratKenlinci[i] > max {
			max = beratKenlinci[i]
		}
	}
	
	fmt.Printf("\nBerat kelinci terkecil: %.2f kg\n", max)
	fmt.Printf("\nBerat kelinci terbesar: %.2f kg\n", min)
}