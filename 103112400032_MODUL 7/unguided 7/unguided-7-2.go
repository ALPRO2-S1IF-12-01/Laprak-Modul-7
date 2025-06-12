// daffa tsaqifna f
package main

import "fmt"

type IkanData struct {
	berat       []float64
	jumlahIkan  int
	isiPerWadah int
}

const maxSize = 1000

func inputIkan() IkanData {
	var x, y int

	fmt.Print("Masukkan jumlah ikan dan kapasitas wadah (x y): ")
	fmt.Scan(&x, &y)

	if x <= 0 || x > maxSize || y <= 0 {
		fmt.Println("Input tidak valid.")
		return IkanData{}
	}

	arr := make([]float64, x)
	fmt.Printf("Masukkan %d berat ikan:\n", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}

	return IkanData{
		berat:       arr,
		jumlahIkan:  x,
		isiPerWadah: y,
	}
}

func hitungTotalPerWadah(data IkanData) []float64 {
	var hasil []float64
	jumlah := data.jumlahIkan
	kaps := data.isiPerWadah

	for i := 0; i < jumlah; i += kaps {
		total := 0.0
		for j := i; j < i+kaps && j < jumlah; j++ {
			total += data.berat[j]
		}
		hasil = append(hasil, total)
	}

	return hasil
}

func hitungRataRata(wadah []float64) float64 {
	if len(wadah) == 0 {
		return 0
	}
	total := 0.0
	for _, val := range wadah {
		total += val
	}
	return total / float64(len(wadah))
}

func main() {
	data := inputIkan()
	if data.jumlahIkan == 0 {
		return
	}

	wadah := hitungTotalPerWadah(data)
	for _, berat := range wadah {
		fmt.Printf("%.2f ", berat)
	}
	fmt.Println()

	rata := hitungRataRata(wadah)
	fmt.Printf("%.2f\n", rata)
}
