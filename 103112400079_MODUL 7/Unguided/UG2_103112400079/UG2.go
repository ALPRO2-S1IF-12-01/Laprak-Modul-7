package main

//Muhammad Faris Rachmadi
//103112400079

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&x, &y)

	var beratIkan [1000]float64

	fmt.Println("Masukkan berat masing-masing ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	fmt.Println("\nTotal berat ikan per wadah:")
	for i := 0; i < x; i += y {
		var total float64
		for j := i; j < i+y && j < x; j++ {
			total += beratIkan[j]
		}
		fmt.Printf("Wadah %d: %.2f kg\n", i/y+1, total)
	}

	fmt.Println("\nRata-rata berat ikan per wadah:")
	for i := 0; i < x; i += y {
		var total float64
		var jumlah int
		for j := i; j < i+y && j < x; j++ {
			total += beratIkan[j]
			jumlah++
		}
		rata := total / float64(jumlah)
		fmt.Printf("Wadah %d: %.2f kg\n", i/y+1, rata)
	}
}
