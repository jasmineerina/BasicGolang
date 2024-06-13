package main

import "fmt"

func cariAngkaBesar(angka []int) int {
	angkaBesar := angka[0]

	for i := 0; i < len(angka); i++ {
		if angkaBesar < angka[i] {
			angkaBesar = angka[i]
		}
	}

	return angkaBesar
}

func cariAngkaKecil(angka []int) int {
	angkaKecil := angka[0]

	for i := 0; i < len(angka); i++ {
		if angkaKecil > angka[i] {
			angkaKecil = angka[i]
		}
	}

	return angkaKecil
}

func main() {
	array := []int{2, 6, 4, 9, 0}

	fmt.Println(cariAngkaBesar(array))
	fmt.Println(cariAngkaKecil(array))
}
