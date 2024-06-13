package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hitungKata(kata string) string {
	jumlahKata := strings.Split(kata, " ")
	s := strconv.Itoa(len(jumlahKata)) // parse int to str

	return "Jumlah kalimat = " + s
}

func hitungHuruf(kata string) string {
	// jumlahHuruf := strings.Split(kata, " ")
	// s := strconv.Itoa(len(jumlahHuruf)) // parse int to str

	// return "Jumlah huruf = " + s

	jumlahTotalHuruf := 0
	jumlahHuruf := strings.Split(kata, " ")

	for i := 0; i < len(jumlahHuruf); i++ {
		jumlahTotalHuruf += len(jumlahHuruf[i])
	}

	s := strconv.Itoa(jumlahTotalHuruf)

	return "jumlah huruf ada = " + s
}

func main() {
	fmt.Println(hitungKata("Halo, nama saya tukul arwana"))
	fmt.Println(hitungHuruf("Halo, nama saya tukul arwana"))
}
