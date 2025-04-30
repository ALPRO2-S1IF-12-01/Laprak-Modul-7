package main

import "fmt"

func main() {
	var n int
	var berat [1000]float64

	fmt.Print("Jumlah anak kelinci: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan berat masing-masing kelinci:")
	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 0; i < n; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
