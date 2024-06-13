package main

import (
	"fmt"
	"strings"
)

func cekDouble(kata string) string {
	kataAkhir := ""
	kata = strings.ToLower(kata) // untuk cek lebih konsisten

	for i := 0; i < len(kata); i++ {
		char := kata[i]
		duplikasi := false

		for j := 0; j < len(kataAkhir); j++ {
			if char == kataAkhir[j] {
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
