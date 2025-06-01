//Nama : Anggun Wahyu Widiyana (103112480280)
package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan banyaknya anak kelinci: ")
    fmt.Scan(&n)

    var berat = make([]float64, n)
    fmt.Println("Masukkan berat anak kelinci:")
    for i := range berat {
        fmt.Scan(&berat[i])
    }

    min := berat[0]
    max := berat[0]

    for _, b := range berat {
        if b < min {
            min = b
        }
        if b > max {
            max = b
        }
    }

    fmt.Printf("Berat terkecil: %.2f\n", min)
    fmt.Printf("Berat terbesar: %.2f\n", max)
}