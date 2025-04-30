//RAJA MUHAMMAD LUFHTI_103112400027
package main

import "fmt"

func main() {
	var jumlahAnak int
	var berat [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahAnak)

	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < jumlahAnak; i++ {
		fmt.Printf("Berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}

	beratTerkecil := berat[0]
	beratTerbesar := berat[0]

	for i := 1; i < jumlahAnak; i++ {
		if berat[i] < beratTerkecil {
			beratTerkecil = berat[i]
		}
		if berat[i] > beratTerbesar {
			beratTerbesar = berat[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f\n", beratTerkecil)
	fmt.Printf("Berat terbesar: %.2f\n", beratTerbesar)
}
