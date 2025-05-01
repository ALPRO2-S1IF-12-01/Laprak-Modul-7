package main

import "fmt"

func main() {
	var n int

	fmt.Print("jumlah anak kelinci : ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("jumlah anak kelinci tidak valid.")
		return
	}

	var beratKelinci [1000]float64
	fmt.Println("berat anak kelinci :")
	for i := 0; i < n; i++ {
		fmt.Printf("berat ke-%d: ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	beratTerkecil := beratKelinci[0]
	beratTerbesar := beratKelinci[0]

	for i := 1; i < n; i++ {
		if beratKelinci[i] < beratTerkecil {
			beratTerkecil = beratKelinci[i]
		}
		if beratKelinci[i] > beratTerbesar {
			beratTerbesar = beratKelinci[i]
		}
	}

	fmt.Printf("berat terkecil: %.2f kg\n", beratTerkecil)
	fmt.Printf("berat terbesar: %.2f kg\n", beratTerbesar)
}
