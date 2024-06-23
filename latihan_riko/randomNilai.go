package main

import (
	"fmt"
	"math/rand"
)

// buat func tanpa parameter
// buat var arr[]int untuk angka random
// lakukan perulangan 5x, melakukan random
// hasilnya disimpan ke var hasil int
// perulangan lagi untuk cek setiap angka masuk ke kondisi skor apa
// nilai disimpan ke array baru yang berisi skor

func tentukanNilai() string {
	var angka []int
	var skor []int

	for i := 0; i < 5; i++ {
		randomizer := rand.Intn(100-1) + 1 // nentuin rentang min hingga maks
		angka = append(angka, randomizer)
	}

	for i := 0; i < len(angka); i++ {
		if angka[i] > 80 {
			fmt.Println(angka[i], "Luar Biasa")
		} else if angka[i] <= 80 && angka[i] > 70 {
			fmt.Println(angka[i], "Baik")
		} else if angka[i] <= 70 && angka[i] > 60 {
			fmt.Println(angka[i], "Cukup")
		} else {
			fmt.Println(angka[i], "Kurang")
		}
	}

	angka = append(angka, skor...)

	return ""
}

func main() {
	fmt.Println(tentukanNilai())
}
