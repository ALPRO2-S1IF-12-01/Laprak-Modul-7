package main

//103112400078 Mohammad Reyhan Aretha Fatin
import "fmt"

func main() {
	var totalIkan, ikanPerWadah int
	var beratIkan [1000]float64

	fmt.Print("Masukkan jumlah ikan (x) dan jumlah ikan per wadah (y): ")
	fmt.Scan(&totalIkan, &ikanPerWadah)

	fmt.Print("Masukkan berat setiap ikan: ")
	for i := 0; i < totalIkan; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (totalIkan + ikanPerWadah - 1) / ikanPerWadah
	totalBeratPerWadah := make([]float64, jumlahWadah)
	jumlahIkanPerWadah := make([]int, jumlahWadah)

	for i := 0; i < totalIkan; i++ {
		wadahIdx := i / ikanPerWadah
		totalBeratPerWadah[wadahIdx] += beratIkan[i]
		jumlahIkanPerWadah[wadahIdx]++
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("| Wadah ke-%d: %.2f kg |", i+1, totalBeratPerWadah[i])
		if jumlahIkanPerWadah[i] > 0 {
			fmt.Printf(" Berat rata-rata: %.2f kg |", totalBeratPerWadah[i]/float64(jumlahIkanPerWadah[i]))
		}
		fmt.Println()
	}
}
