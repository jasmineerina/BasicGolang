package main

import "fmt"

// slice adalah reference elemen array. slice bisa dibuat dan dihasilkan dari manipulasi
// sebuah array atau slice lainnya.
// Slice merupakan Data Reference, perubahan tiap data akan berdampak pd slice lainnya.

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println("--------------------------------")

	fmt.Println(fruits[0]) // [apple]

	fmt.Println("--------------------------------")

	// perbedaan slice & array
	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array

	fmt.Println(fruitsA, fruitsB, fruitsC)

	fmt.Println("--------------------------------")

	//
	var newFruits = fruits[0:2] // akses slice fruits dimulai dr indeks 0 hingga sebelum indeks ke 2.
	fmt.Println(newFruits)      // output: apple grape

	fmt.Println("--------------------------------")

	// slice merupakan tipe data reference

	var aFruits = fruits[0:3] // apple grape banana melon
	var bFruits = fruits[1:4] // apple grape banana

	var aaFruits = aFruits[1:2] // grape
	var baFruits = bFruits[0:1] // grape

	fmt.Println(aFruits, bFruits, aaFruits, baFruits)

	fmt.Println("--------------------------------")

	baFruits[0] = "pinnaple" // mengubah indeks[0] dari grape menjadi pinnaple

	fmt.Println(aFruits, bFruits, aaFruits, baFruits) // semua indeks ke 0 menjadi pinnaple

}
