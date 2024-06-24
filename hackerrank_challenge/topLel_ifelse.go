package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	//     buat variable input
	// var kalimat string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	kalimat := scanner.Text()

	//     ubah kalimat menjadi array kata
	kata := strings.Split(kalimat, " ")

	//     buat variable untuk simpan skor
	var skor int

	//     mencari poin dari perkata yang muncul
	//     looping dulu untuk cek per kata
	for i := 0; i < len(kata); i++ {
		if kata[i] == "lol" {
			skor += 1
		} else if kata[i] == "rofl" {
			skor += 2
		} else if kata[i] == "lmao" {
			skor += 3
		} else if kata[i] == "lel" {
			skor += 4
		}
	}

	//     tentuin laugh points
	var hasil string

	// if skor <= 5 && skor > 1 {
	//     hasil = "Patient has bright red face"
	// } else if skor <= 12 && skor > 6 {
	//      hasil ="Patient is unable to speak"
	// } else if skor <= 20 && skor > 13 {
	//      hasil ="Patient's sides are mildly bruised"
	// } else if skor <= 30 && skor > 21 {
	//      hasil = "Patient has broken jaw, fractured ribs"
	// } else if skor <= 49 && skor > 31 {
	//      hasil = "Patient is in a coma"
	// } else if skor > 50 {
	//      hasil = "Patient is dead"
	// }

	if skor > 50 {
		hasil = "Patient is dead"
	} else if skor >= 31 {
		hasil = "Patient is in coma"
	} else if skor >= 21 {
		hasil = "Patient has broken jaw, fractured ribs"
	} else if skor >= 13 {
		hasil = "Patient's sides are mildly bruised"
	} else if skor >= 6 {
		hasil = "Patient is unable to speak"
	} else if skor >= 1 {
		hasil = "Patient has bright red face"
	}

	fmt.Println(hasil)

}
