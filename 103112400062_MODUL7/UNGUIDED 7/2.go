//M. DAVI ILYAS RENALDO_103112400062
package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("masukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scanln(&x)
	fmt.Print("masukkan jumlah wadah (y): ")
	fmt.Scanln(&y)

	var totalBerat [1000]float64

	for i := 0; i < y; i++ {
		fmt.Printf("masukkan berat ikan untuk wadah ke-%d (jumlah ikan: %d): ", i+1, x)
		for j := 0; j < x; j++ {
			var berat float64
			fmt.Scanln(&berat)
			totalBerat[i] += berat
		}
	}

	for i := 0; i < y; i++ {
		fmt.Printf("total berat ikan pada wadah ke-%d: %.2f kg\n", i+1, totalBerat[i])
	}

	var totalRataRata float64
	for i := 0; i < y; i++ {
		totalRataRata += totalBerat[i]
	}

	fmt.Printf("rata-rata berat ikan di setiap wadah: %.2f kg\n", totalRataRata/float64(y))
}