package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("jumlah ikan dan jumlah ikan per wadah : ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 {
		fmt.Println("tidak valid.")
		return
	}

	var beratIkan [1000]float64

	fmt.Println("berat ikan :")
	for i := 0; i < x; i++ {
		fmt.Printf("berat ikan ke-%d: ", i+1)
		fmt.Scan(&beratIkan[i])
	}

	var beratPerWadah []float64
	var totalBerat float64
	for i := 0; i < x; i++ {
		totalBerat += beratIkan[i]

		if (i+1)%y == 0 || i == x-1 {
			beratPerWadah = append(beratPerWadah, totalBerat)
			totalBerat = 0
		}
	}

	fmt.Println("berat ikan di setiap wadah :")
	for i, berat := range beratPerWadah {
		fmt.Printf("Wadah %d: %.2f kg\n", i+1, berat)
	}

	rataRata := beratPerWadah[0] / float64(y)
	fmt.Printf("\nrata-rata berat ikan per wadah: %.2f kg\n", rataRata)
}
