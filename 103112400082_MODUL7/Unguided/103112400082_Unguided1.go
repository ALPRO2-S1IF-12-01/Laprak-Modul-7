//Rizkina Azizah 103112400082
package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64
	fmt.Print("jumlah anak kelinci: ")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Printf("berat anak kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat[i])
	}

	min, max := berat[0], berat[0]
	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Println("Berat kelinci terkecil:", min)
	fmt.Println("Berat kelinci terbesar:", max)
}