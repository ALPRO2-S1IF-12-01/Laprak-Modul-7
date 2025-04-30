package main

//103112400049 Hisyam Nurdiatmoko

import "fmt"

func main() {
	var hisyam int
	var berat [1000]float64
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&hisyam)
	for i := 0; i < hisyam; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&berat[i])
	}
	min := berat[0]
	max := berat[0]
	for i := 1; i < hisyam; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}
	fmt.Printf("Berat kelinci terkecil: %.2f\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", max)
}
