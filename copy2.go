package main

import "fmt"

func main() {
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinneaple"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon, pinneaple potato
	fmt.Println(src) // watermelon pinneaple
	fmt.Println(n)   // 2

	// dst masih tetap 3 element, tapi 2 element pertama sama seperti src
	// element terakhir dst isinya tidak berubah tetap "potato"
	// hal ini karena proses copy() hanya memutasi elemen ke-1 dan ke-2 milik dst
	// karena pada src hanya dua itu elemennya

}
