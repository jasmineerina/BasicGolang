package main

import "fmt"

func hitungHarga(hargaProduk []float64) (float64, float64, float64, string) {
	totalHarga := 0.0

	for _, harga := range hargaProduk {
		totalHarga += harga
	}

	diskon := 0.0
	hadiah := ""

	if totalHarga > 400000 {
		diskon = 0.01
		hadiah = "Ransel"
	} else if totalHarga > 200000 {
		diskon = 0.07
		hadiah = "Payung"
	} else if totalHarga > 70000 {
		diskon = 0.05
		hadiah = "Topi"
	}

	totalDiskon := totalHarga * diskon
	totalAkhir := totalHarga - totalDiskon

	return totalHarga, totalDiskon, totalAkhir, hadiah
}

func main() {
	hargaProduk := []float64{2000, 50000, 100000}

	totalHarga, totalDiskon, totalAkhir, hadiah := hitungHarga(hargaProduk)

	fmt.Println(hitungHarga(hargaProduk))
	fmt.Println("Total harga =", totalHarga)
	fmt.Println("Diskon =", totalDiskon)
	fmt.Println("Total akhir =", totalAkhir)
	fmt.Println("Bonus yang didapat =", hadiah)

}
