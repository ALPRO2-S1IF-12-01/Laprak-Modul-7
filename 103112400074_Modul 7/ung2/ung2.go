// Ahmad Ruba'i
// 103112400074
package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Print("Jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Kapasitas per wadah: ")
	fmt.Scan(&y)

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	var total float64
	var totalPerWadah []float64

	for i := 0; i < x; i += y {
		total = 0
		for j := i; j < i+y && j < x; j++ {
			total += berat[j]
		}
		totalPerWadah = append(totalPerWadah, total)
	}

	for i, t := range totalPerWadah {
		fmt.Printf("Wadah %d: %.1f\n", i+1, t)
	}

	var jumlah float64
	for _, t := range totalPerWadah {
		jumlah += t
	}
	rata := jumlah / float64(len(totalPerWadah))
	fmt.Printf("Rata-rata: %.1f\n", rata)
}
