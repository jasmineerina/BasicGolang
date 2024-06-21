package main

import "fmt"

func main() {

	// fungsi append() digunakan untuk menambahkan elemen pada slice.
	// dan ditaruh di indeks terakhir dr sebuah slice

	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // apple grape banana
	fmt.Println(cFruits) // apple grape banana papaya

	fmt.Println("----------------------------------------")

	var bFruits = fruits[0:2]

	fmt.Println("cap: ", cap(bFruits)) // cap: 3
	fmt.Println("len: ", len(bFruits)) // len: 2

	fmt.Println("----------------------------------------")

	fmt.Println(fruits)  // apple grape banana
	fmt.Println(bFruits) // apple grape

	fmt.Println("----------------------------------------")

	var cbFruits = append(bFruits, "papaya")

	fmt.Println(fruits)   // apple grape papaya
	fmt.Println(bFruits)  // apple grape
	fmt.Println(cbFruits) // apple grape papaya
}
