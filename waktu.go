package main

import "fmt"

func ubahWaktu(waktu int) string {
	// detik ke jam = 3600 detik / 1 jam
	// sisa detik jam = waktu%jam
	// detik ke menit = 60 detik / 1 menit
	// sisa detik menit = waktu%menit

	// jam := 0
	// menit := 0
	// detik := 0

	jam := waktu / 3600
	menit := (waktu % 3600) / 60
	detik := waktu % 60

	fmt.Println(jam, "jam ", menit, "menit ", detik, "detik ")

	return ""
}

func main() {
	fmt.Println(ubahWaktu(11453))
}
