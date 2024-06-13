package main

import "fmt"

func hitungTinggi(tinggiAwal float32, lamaHari int, tambahTinggi float32) float32 {
	// 100, 3, 0.5
	tinggiAkhir := tinggiAwal

	for i := 0; i < lamaHari; i++ {
		tinggiAkhir = tinggiAkhir + (tinggiAkhir * tambahTinggi)
	}

	return tinggiAkhir
}

func main() {
	fmt.Println(hitungTinggi(200, 5, 0.05))
	fmt.Println(hitungTinggi(250, 2, 0.02))
}
