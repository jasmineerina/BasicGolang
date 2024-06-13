package main

import "fmt"

func hitungHarga(harga []int) int {
	total := 0

	for _, i := range harga {
		total += i
	}

	// Initialize variables for discount, bonus, and total bonus
	var diskon float64
	var hadiah string
	totalBonus := 0

	// Determine the discount and bonus based on the total price
	if total > 400000 {
		diskon = 0.01
		hadiah = "Ransel"
	} else if total > 200000 {
		diskon = 0.07
		hadiah = "Payung"
	} else if total > 70000 {
		diskon = 0.05
		hadiah = "Topi"
	}

	// Calculate the total bonus
	totalBonus = int(float64(total) * diskon)

	// Calculate the final price after applying the discount
	totalAkhir := total - totalBonus

	return totalAkhir
}

func main() {
	hargaProduk := []int{2000, 50000, 150000}

	totalHarga := hitungHarga(hargaProduk)
	fmt.Println("Total harga : ", totalHarga)
}
