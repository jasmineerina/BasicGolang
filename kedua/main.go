package main

// ini kita belajar tentang tipe data dan nilai default
// kalau sebuah variable tidak diisi maka akan print " "
// string = "halo", "dst"
// int = 1, 2, 3 (bil. bulat)
//  float64 = 3.14, 100.1 (desimal)
// bool = true false

import (
	"fmt"
)

func main() {
	var cobadulu string = "ayo lagi kita belajar go"
	var ayo string
	ayo = "jangan lupa minum air putih"

	var belajar bool = true

	// tanpa pakai var
	nama := "jasmine"
	angka := 15052001

	// := berfungsi buat membuat var baru, jadi gabisa dibuat kayak dibawah
	// age := 12
	// age := 13
	// code spt di atas tdk bisa di run

	fmt.Println(cobadulu)
	fmt.Println(ayo)
	fmt.Println(belajar, angka)
	fmt.Println(nama)

}
