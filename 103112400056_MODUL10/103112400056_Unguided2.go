// Felix Pedrosa V

package main

import "fmt"

func main() {
	const maxIkan = 1000
	var beratIkan [maxIkan]float64
	var x, y int

	// Input jumlah ikan dan jumlah ikan per wadah
	fmt.Print("Masukkan jumlah ikan yang akan dijual (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jumlah ikan per wadah (y): ")
	fmt.Scan(&y)

	// Input berat ikan
	fmt.Println("Masukkan berat ikan satu per satu:")
	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	// Proses pengelompokan ke dalam wadah
	var wadah [][]float64
	for i := 0; i < x; i += y {
		akhir := i + y
		if akhir > x {
			akhir = x
		}
		wadah = append(wadah, beratIkan[i:akhir])
	}

	// Hitung dan tampilkan total berat setiap wadah
	var totalSemuaWadah float64
	fmt.Println("\nTotal berat ikan di setiap wadah:")
	for i, w := range wadah {
		var total float64
		for _, berat := range w {
			total += berat
		}
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, total)
		totalSemuaWadah += total
	}

	// Hitung dan tampilkan rata-rata berat per wadah
	rataRata := totalSemuaWadah / float64(len(wadah))
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah: %.2f kg\n", rataRata)
}