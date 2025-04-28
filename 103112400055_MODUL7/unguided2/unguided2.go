//Feros Pedrosa

package main

import "fmt"

func tarif(berat []float64, y int) ([]float64, float64) {
	var total []float64
	var jumlah float64
	count := 0

	for i, b := range berat {
		jumlah += b
		count++
		if count == y || i == len(berat)-1 {
			total = append(total, jumlah)
			jumlah = 0
			count = 0
		}
	}

	var sum float64
	for _, t := range total {
		sum += t
	}

	avg := sum / float64(len(total))
	return total, avg
}

func main() {
	var x, y int
	fmt.Println("Masukkan banyak ikan dan kapasitas wadah:")
	fmt.Scan(&x, &y)

	berat := make([]float64, x)
	fmt.Println("Masukkan berat ikan:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	total, avg := tarif(berat, y)
	fmt.Println("\nTotal berat ikan per wadah:")
	for i, t := range total {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, t)
	}

	fmt.Printf("Rata-rata berat di setiap wadah: %.2f kg\n", avg)
}
