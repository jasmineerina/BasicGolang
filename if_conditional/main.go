package main

import "fmt"

func main() {
	// age := 17

	// if age > 20 {
	// 	fmt.Println("Umur", age, "boleh melakukan Tugas Akhir")
	// } else {
	// 	fmt.Println("Umur", age, "belum boleh mengambil Tugas Akhir")
	// }

	// score := 25
	// var grade string

	// if score <= 50 {
	// 	grade = "E"
	// } else if score <= 60 {
	// 	grade = "D"
	// } else if score <= 70 {
	// 	grade = "C"
	// } else {
	// 	grade = "NULL"
	// }

	// fmt.Println(grade)

	nilai := 78
	var grade string

	if nilai <= 50 {
		fmt.Println("Nilai ", nilai, " tidak mencapai KKM")
	} else {
		fmt.Println("Nilai ", nilai, " mencapai KKM")
	}

	fmt.Println("===========")

	if nilai == 100 {
		grade = "A"
	} else if nilai > 80 {
		grade = "B"
	} else if nilai > 70 {
		grade = "C"
	} else {
		grade = "D"
	}

	fmt.Println("Nilai Tono ", nilai, "maka grade Tono adalah", grade)

}

// if
// if else
// else if
// if bersarang

// maksimal if didalam if cuma boleh 2
// if {
// 	if{

// 	}
// }
