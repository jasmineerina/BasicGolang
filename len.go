package main

import "fmt"

func main() {

	// fungsi len() menghitung jumlah dari sebuah elemen slice

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits))

	fmt.Println("---------------------------------------------")

	// fungsi cap() digunakan untuk menghitung lebar /  kapasitas maks. slice
	// sama seperti len() tetapi bisa berubah seiring operasi slice yang dilakukan

	fmt.Println(cap(fruits)) // cap: 4

	fmt.Println("fruits[0:3] :")
	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // len: 4

	fmt.Println("fruits[1:4] :")
	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3

}
