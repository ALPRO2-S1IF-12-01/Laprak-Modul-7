package main

//Muhammad Faris Rachmadi
//103112400079

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("masukkan jumlah kelinci :")
	fmt.Scan(&n)

	var beratkelinci [1000]float64
	for i := 0; i < n; i++ {
		fmt.Printf("masukkan berat kelinci ke-%d (kg):", i+1)
		fmt.Scan(&beratkelinci[i])
	}
	min, max := beratkelinci[0], beratkelinci[0]
	for i := 1; i < n; i++ {
		if beratkelinci[i] > min {
			min = beratkelinci[i]
		}
		if beratkelinci[i] < max {
			max = beratkelinci[i]
		}
	}
	fmt.Printf("\nberat kelinci terkecil : %.2f kg\n", max)
	fmt.Printf("\nberat kelinci terbesar : %.2f kg\n", min)

}
