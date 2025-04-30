package main

import "fmt"

func main() {
	var jumlahKelinci int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Println("masukkan jumlah kelinci: ")
	fmt.Scan(&jumlahKelinci)

	if jumlahKelinci <= 0 {
		return
	}
	daftarBerat := make([]float64, jumlahKelinci)

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("berat kelinci ke-%d: ", i+1)
		fmt.Scan(&daftarBerat[i])
	}

	palingRingan := daftarBerat[0]
	palingBerat := daftarBerat[0]

	for i := 1; i < jumlahKelinci; i++ {
		beratKelinciIni := daftarBerat[i]

		if beratKelinciIni < palingRingan {
			palingRingan = beratKelinciIni
		}

		if beratKelinciIni > palingBerat {
			palingBerat = beratKelinciIni
		}
	}

	fmt.Println("Kelinci paling ringan:", palingRingan)
	fmt.Println("Kelinci paling berat:", palingBerat)
}