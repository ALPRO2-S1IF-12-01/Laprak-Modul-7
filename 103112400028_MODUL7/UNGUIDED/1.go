package main

import "fmt"
	
func main() {
	var n int
	fmt.Print("Mausukan jumlah tanaman")
	fmt.Scan(&n)

	var tinggiTanaman [500]float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan tinggi tanaman ke-%d (cm): ", i+i)
		fmt.Scan(&tinggiTanaman[i])																																																																																																																
	}

	min, max := tinggiTanaman[0], tinggiTanaman[0]
	for i := 1; i < n; i++ {
		if tinggiTanaman[i] < min {
			min = tinggiTanaman[i]
		}
	} 												
	
	fmt.Printf("\nTinggi tanaman tertinggi: %.2f cm\n", max)
	fmt.Printf("Tingi tanaman terpendek: %.2f cm\n", min)
																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																																			
}																																																																										