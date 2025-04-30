package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan jumlah tanaman: ")
	fmt.Scan(&n)
	var tinggitanaman [500]float64
	for i := 0; i < n; i++ {
		fmt.Printf("masukan tinggi tanaman ke -%d(cm) : ", i+1)
		fmt.Scan(&tinggitanaman[i])
	}
	min, max := tinggitanaman[0], tinggitanaman[0]
	for i := 1; i < n; i++ {
		if tinggitanaman[i] < min {
			min = tinggitanaman[i]
		}
		if tinggitanaman[i] > max {
			max = tinggitanaman[i]
		}
	}
	fmt.Printf("\nTinggi Tanaman Tertinggi: %.2f cm", max)
	fmt.Printf("\nTinggi tanaman Terpendek: %.2f cm", min)
}
