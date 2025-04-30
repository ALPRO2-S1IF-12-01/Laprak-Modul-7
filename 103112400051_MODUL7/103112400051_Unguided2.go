package main

import "fmt"

func main() {
	var a, b int
	var data [1000]float64
	var hasil [1000]float64

	fmt.Print("Masukkan jumlah ikan dan wadah: ")
	fmt.Scan(&a, &b)

	fmt.Println("Berat ikannya:")
	for i := 0; i < a; i++ {
		fmt.Scan(&data[i])
	}

	for i := 0; i < a; i++ {
		pos := i % b
		hasil[pos] = hasil[pos] + data[i]
	}

	fmt.Println("Isi wadah:")
	for i := 0; i < b; i++ {
		fmt.Print(hasil[i], " ")
	}
	fmt.Println()

	var jum float64 = 0
	for i := 0; i < b; i++ {
		jum = jum + hasil[i]
	}
	rata := jum / float64(b)
	fmt.Println("Rata-rata:", rata)
}
