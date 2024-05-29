package main

import "fmt"

// kalau operasi &&, pasangan false hasilnya bernilai false
// operasi ||, kalau kedua pasangan adalah false hasilnya akan false
// operasi !, akan mengubah nilai.

func main() {
	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println("Mahasiswa Lulus = ", lulus)
}
