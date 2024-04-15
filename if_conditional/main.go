package main

import "fmt"

func main() {
	// age := 17

	// if age > 20 {
	// 	fmt.Println("Umur", age, "boleh melakukan Tugas Akhir")
	// } else {
	// 	fmt.Println("Umur", age, "belum boleh mengambil Tugas Akhir")
	// }

	score := 25
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
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
