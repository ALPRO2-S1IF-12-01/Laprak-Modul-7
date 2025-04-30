package main
//103112400078 Mohammad Reyhan Aretha Fatin
import "fmt"

func hitungMinMaxRata(arr []float64, n int) (min, max, rata float64) {
	min, max = arr[0], arr[0]
	var total, count float64

	for i := 0; i < n; i++ {
		if arr[i] > 0 {
			if arr[i] < min {
				min = arr[i]
			}
			if arr[i] > max {
				max = arr[i]
			}
			total += arr[i]
			count++
		}
	}
	if count > 0 {
		rata = total / count
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

	var dataBalita [100]float64
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	min, max, rata := hitungMinMaxRata(dataBalita[:n], n)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rata-rata berat balita: %.2f kg\n", rata)
}
