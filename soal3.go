package main

import "fmt"

func cekDouble(kata string) string {
	kataAkhir := ""
	duplikasi := false

	for i := 0; i < len(kata); i++ {
		char := kata[i]
		duplikasi = false

		for j := 0; j < len(kata); j++ {
			if char == kata[j] {
				duplikasi = true
				break
			}
		}

		if !duplikasi {
			kataAkhir += string(char)
		}
	}

	return kataAkhir
}

func main() {

	fmt.Println(cekDouble("Imagination"))
	fmt.Println(cekDouble("Association"))

}
