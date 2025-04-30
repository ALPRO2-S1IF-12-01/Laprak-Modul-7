package main
// 103112400084 NUFAIL ALAUDDIN TSAQIF
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

	fmt.Println("\nTotal berat tiap wadah:")
	for i, berat := range totalBeratPerWadah {
		fmt.Printf("| Wadah ke-%d: %.2f kg |", i+1, berat)
	}
	fmt.Println()

	fmt.Println("\nRata-rata berat tiap wadah:")
	for i := 0; i < jumlahWadah; i++ {
		if jumlahIkanPerWadah[i] > 0 {
			rataRata := totalBeratPerWadah[i] / float64(jumlahIkanPerWadah[i])
			fmt.Printf("| Berat rata-rata wadah ke-%d: %.2f kg |", i+1, rataRata)
		}
	}
	fmt.Println()
}