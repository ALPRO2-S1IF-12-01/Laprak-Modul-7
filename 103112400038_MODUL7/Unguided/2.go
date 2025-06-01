//Nama	: Hakan Ismail Afnan 
//NIM	: 103112400038
package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    // Baca data berat ikan ke array
    var ikan [1000]float64
    for i := 0; i < x; i++ {
        fmt.Scan(&ikan[i])
    }

    // Hitung total per wadah
    var totalWadah [1000]float64
    jumlahWadah := 0
    totalSemua := 0.0
    for i := 0; i < x; i += y {
        akhir := i + y
        if akhir > x {
            akhir = x
        }
        total := 0.0
        for j := i; j < akhir; j++ {
            total += ikan[j]
        }
        totalWadah[jumlahWadah] = total
        jumlahWadah++
        totalSemua += total
    }

    // Cetak total per wadah
    for i := 0; i < jumlahWadah; i++ {
        fmt.Printf("%.2f ", totalWadah[i])
    }
    fmt.Println()

    // Cetak rata-rata per ikan
    rata := totalSemua / float64(x)
    fmt.Printf("Rata-rata: %.2f\n", rata)
}