package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah tanaman: ")
	fmt.Scan(&n)

	var tinggiTanaman [500] float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan tinggi tanaman ke-%d (cm): ", i+1)
		fmt.Scan(&tinggiTanaman[i])
	}

	min, max := tinggiTanaman[0], tinggiTanaman[0]
	for i := 1; i < n; i++ {
		if tinggiTanaman[1] < min {
			min = tinggiTanaman[1]
		}
		if tinggiTanaman[1] > max {
			max = tinggiTanaman[1]
		}
	}

	fmt.Printf("\nTinggi tanaman tertinggi: %.2f cm\n", max)
	fmt.Printf("Tinggi tanaman terpendek: %.2f cm\n", min)
}