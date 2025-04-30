// JESIKA METANIA RAHMA ARIFIN
// 10311240080

package main

import "fmt"

func main() {
	var anak_Kelinci int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&anak_Kelinci)

	if anak_Kelinci> 1000 {
		fmt.Println("Jumlah anak kelinci melebihi kapasitas (maksimal 1000).")
		return
	}
	var berat [1000]float64
	for i := 0; i < anak_Kelinci; i++ {
		fmt.Printf("Berat kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	min := berat[0]
	max := berat[0]

	for i := 1; i < anak_Kelinci; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	
	fmt.Printf("Berat terkecil: %.2f\n", min)
	fmt.Printf("Berat terbesar: %.2f\n", max)
}