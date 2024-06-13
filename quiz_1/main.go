package main

import "fmt"

func main() {
	nilaiUjian := [4]int{70, 60, 90, 87}

	hasilNilai := 0
	for _, nilai := range nilaiUjian {
		hasilNilai += nilai
	}
	fmt.Println("total nilai :", hasilNilai)

	hasilNilai /= len(nilaiUjian)
	fmt.Println("rata rata nilai :", hasilNilai)

	var grade string
	if hasilNilai == 100 {
		grade = "A"
	} else if hasilNilai >= 85 {
		grade = "B"
	} else if hasilNilai >= 75 {
		grade = "C"
	} else {
		grade = "TIDAK LULUS"
	}

	fmt.Println("Nilai Tono dinyatakan ", grade)
}
