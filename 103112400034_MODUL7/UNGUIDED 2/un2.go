package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64

	fmt.Print("Jumlah ikan: ")
	fmt.Scan(&x)

	fmt.Print("Kapasitas per wadah: ")
	fmt.Scan(&y)

	if x <= 0 || x > 1000 || y <= 0 {
		fmt.Println("Input nggak valid")
		return
	}

	for i := 0; i < x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	jumlahWadah := (x + y - 1) / y
	total := make([]float64, jumlahWadah)

	for i := 0; i < x; i++ {
		wadah := i / y
		total[wadah] += berat[i]
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("Wadah %d: %g\n", i+1, total[i])
	}

	var jumlah float64
	for i := 0; i < jumlahWadah; i++ {
		jumlah += total[i]
	}

	fmt.Printf("Rata-rata: %g\n", jumlah/float64(jumlahWadah))
}
