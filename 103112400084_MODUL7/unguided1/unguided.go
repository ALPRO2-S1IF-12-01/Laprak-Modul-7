package main
// 103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"

func main() {
	var jumlahKelinci int
	var beratKelinci [1000]float64

	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&jumlahKelinci)

	for i := 0; i < jumlahKelinci; i++ {
		fmt.Printf("Masukkan berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&beratKelinci[i])
	}

	min, max := beratKelinci[0], beratKelinci[0]

	for i := 1; i < jumlahKelinci; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	fmt.Printf("Berat kelinci terkecil: %.2f kg\n", min)
	fmt.Printf("Berat kelinci terbesar: %.2f kg\n", max)
}
