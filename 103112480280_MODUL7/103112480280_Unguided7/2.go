//Nama : Anggun Wahyu Widiyana (103112480280)
package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)

    // Baca data berat ikan
    ikan := make([]float64, x)
    for i := 0; i < x; i++ {
        fmt.Scan(&ikan[i])
    }

    // Hitung total per wadah
    var totalWadah []float64
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
        totalWadah = append(totalWadah, total)
        totalSemua += total
    }

    // Cetak total per wadah
    for _, t := range totalWadah {
        fmt.Printf("%.2f ", t)
    }
    fmt.Println()

    // Cetak rata-rata per ikan
    rata := totalSemua / float64(x)
    fmt.Printf("Rata-rata: %.2f\n", rata)
}