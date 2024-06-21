package main

import (
	"fmt"
)

// fungsi copy() untuk copy elements slice pada 'src' (param ke 2),
//  'dst' (param ke pertama)

func main() {

	// 'dst' dipersiapkan dengan lebar 3 elemen
	// slice 'src' isinya 4 elemen di copy ke dst
	// isi 'dst' ada 3 elemen yang sama spt 'src', hasil dari copy()

	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3

	// yang tercopy hanya 3 (walau 'src' ada 4 elemen)
	// hal ini karena copy() hanya copy elemen sebanyak len(dst)

}
