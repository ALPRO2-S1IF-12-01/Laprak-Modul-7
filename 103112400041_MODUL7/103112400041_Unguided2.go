//BERTHA ADELA
//103112400041
package main

import "fmt"

func main() {
	var dataBerat[1000] float64
	var ikanAkanDiJual, ikanTiapWadah int
	var maxWadah[1000] float64
	fmt.Print("Berapa ikan yang mau dijual dan berapa ikan tiap wadahnya: ")
	fmt.Scanln(&ikanAkanDiJual, &ikanTiapWadah)
	if ikanAkanDiJual <= len(dataBerat) {
		wadah := (ikanAkanDiJual+ikanTiapWadah-1)/ikanTiapWadah 
		for i := 0; i < ikanAkanDiJual; i++ {
				fmt.Printf("Berat ikan (%d/%d): ", i+1, ikanAkanDiJual)
				fmt.Scanln(&dataBerat[i])
		}
		for i := 0; i < wadah; i++ {
			mulai := i * ikanTiapWadah
			akhir := mulai + ikanTiapWadah
			if akhir > ikanAkanDiJual {
				akhir = ikanAkanDiJual
			}
			beratTiapWadah := 0.0
			for j := mulai; j < akhir; j++ {
				beratTiapWadah += dataBerat[j]
			}
			maxWadah[i] = beratTiapWadah
		}
		fmt.Println("Total berat per wadah:")
		for i := 0; i < wadah; i++ {
			fmt.Printf("%.2f ", maxWadah[i])
		}
		fmt.Println()
		total := 0.0
		for i := 0; i < wadah; i++ {
			total += maxWadah[i]
		}
		rataRata := total / float64(wadah)
		fmt.Printf("Rata-rata berat setiap wadah: %.2f\n", rataRata)
	} else {
		fmt.Println("Jumlah ikan melebihi batas maksimal(1000)")
	}
	
}