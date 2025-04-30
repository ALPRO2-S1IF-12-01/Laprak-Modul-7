package main
// 103112400084 NUFAIL ALAUDDIN TSAQIF
import "fmt"

func hitungMinMaxRata(arr []float64) (min, max, rata float64) {
	if len(arr) == 0 {
		return 0, 0, 0
	}

	min, max = arr[0], arr[0]
	var total float64

	for _, v := range arr {
		if v > 0 { 
			if v < min {
				min = v
			}
			if v > max {
				max = v
			}
			total += v
		}
	}

	if len(arr) > 0 {
		rata = total / float64(len(arr))
	}
	return
}

func main() {
	var n int
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data tidak valid.")
		return
	}

	var dataBalita = make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	min, max, rata := hitungMinMaxRata(dataBalita)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
