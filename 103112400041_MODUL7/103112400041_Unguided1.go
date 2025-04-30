//BERTHA ADELA
//103112400041
package main

import "fmt"

func main() {
	var banyakAnakKelinci int
	var beratKelinci float64
	var dataBerat[1000]float64
	fmt.Print("Banyak anak kelinci: ")
	fmt.Scanln(&banyakAnakKelinci)
	beratMin := dataBerat[0]
	beratMax := dataBerat[0]
	hitung := 0
	if banyakAnakKelinci <= len(dataBerat) {
		for i := 0; i < banyakAnakKelinci; i++ {
			hitung += 1
			fmt.Printf("Masukkan berat anak kelinci (%d/%d): ", hitung, banyakAnakKelinci)
			fmt.Scanln(&beratKelinci)
			dataBerat[i] = beratKelinci
			if i == 0 {
				beratMin = beratKelinci
				beratMax = beratKelinci
			} else {
				if beratKelinci > beratMax {
					beratMax = beratKelinci
				}
				if beratKelinci < beratMin {
					beratMin = beratKelinci
				}
			}
		}
		fmt.Printf("Berat anak kelinci terkecilnya adalah %.2f", beratMin)
		fmt.Printf("\nBerat anak kelinci terbesarnya adalah %.2f", beratMax)
	} else {
		fmt.Println("Jumlah anak kelinci melebihi batas maksimal (1000)")
	}
}