package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukkan jumlah tanaman :")
	fmt.Scan(&n)

	var tinggitanaman [500]float64
	for i := 0; i < n; i++ {
		fmt.Printf("masukkan tinggi tanaman ke-%d (cm):", i+1)
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
	fmt.Printf("\ntinggi tanaman tertinggi : %.2f cm\n", max)
	fmt.Printf("\ntinggi tanaman terpendek : %.2f cm\n", min)

}
