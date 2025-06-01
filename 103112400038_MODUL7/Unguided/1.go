//Nama	: Hakan Ismail Afnan 
//NIM	: 103112400038
package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan banyaknya anak kelinci: ")
    fmt.Scan(&n)

    var berat [1000]float64
    fmt.Println("Masukkan berat anak kelinci:")
    for i := 0; i < n; i++ {
        fmt.Scan(&berat[i])
    }

    min := berat[0]
    max := berat[0]

    for i := 1; i < n; i++ {
        if berat[i] < min {
            min = berat[i]
        }
        if berat[i] > max {
            max = berat[i]
        }
    }

    fmt.Printf("Berat terkecil: %.2f\n", min)
    fmt.Printf("Berat terbesar: %.2f\n", max)
}