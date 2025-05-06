//Muhammad Zaky Mubarok
package main

import (
	"fmt"
)

func hitungMinMax(daftarBerat []float64) (float64, float64) {
	if len(daftarBerat) == 0 {
		return 0, 0
	}
	beratTerkecil := daftarBerat[0]
	beratTerbesar := daftarBerat[0]

	for _, berat := range daftarBerat[1:] {
		if berat < beratTerkecil {
			beratTerkecil = berat
		}
		if berat > beratTerbesar {
			beratTerbesar = berat
		}
	}
	return beratTerkecil, beratTerbesar
}

func main() {
	var jumlahKelinci int
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	daftarBerat := make([]float64, jumlahKelinci)

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Masukkan berat anak kelinci ke-%d: ", i+1)
		fmt.Scan(&daftarBerat[i])
	}

	beratTerkecil, beratTerbesar := hitungMinMax(daftarBerat)

	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", beratTerbesar)
}
