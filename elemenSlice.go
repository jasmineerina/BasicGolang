package main

import "fmt"

// 3 index adalah teknik slicing untuk pengaksesan elemen
// yang sekaligus menentukan kapasitasnya.

func main() {
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2] // dari indeks 0 sampai sebelum 2, isi elemen 2

	fmt.Println(fruits)      // apple grape banana
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println("--------------------------------------------")

	fmt.Println(aFruits)      // apple grape
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println("--------------------------------------------")

	fmt.Println(bFruits)      // apple grape
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // len: 3

}
