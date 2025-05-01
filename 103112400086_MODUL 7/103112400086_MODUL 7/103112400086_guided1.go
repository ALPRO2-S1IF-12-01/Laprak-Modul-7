package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan jumlah tanaman : ")
	fmt.Scan(&n)

	var tinggiTanaman [500]float64
	for i := 0; i < n; i++ {
		fmt.Printf("masukkan tinggi tanaman ke-%d (cm) : ", i+1)
		fmt.Scan(&tinggiTanaman[i])
	}

	min, max := tinggiTanaman[0], tinggiTanaman[0]
	for i := 0; i < n; i++ {
		if tinggiTanaman[i] < min {
			min = tinggiTanaman[i]
		}
		if tinggiTanaman[i] > max {
			max = tinggiTanaman[i]
		}
	}

	fmt.Printf("\n tinggi tanaman tertiggi : %.2f cm\n", max)
	fmt.Printf("\n tinggi tanaman terpendek : %.2f cm\n", min)
}
