package main

import "fmt"

func cariBilangan(angka []int, kondisi string) string {
	var ganjil []int
	var genap []int

	for i := 0; i < len(angka); i++ {
		if angka[i]%2 == 0 {
			genap = append(genap, angka[i])
		} else {
			ganjil = append(ganjil, angka[i])
		}
	}

	if kondisi == "genap" {
		fmt.Println("Genap: ", genap)
	} else {
		fmt.Println("Ganjil: ", ganjil)
	}

	return ""

}

func main() {
	angka := []int{3, 7, 8, 15, 43, 34}

	fmt.Println(cariBilangan(angka, "genap"))
	fmt.Println(cariBilangan(angka, "ganjil"))
}
