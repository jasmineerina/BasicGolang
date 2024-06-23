package main

import "fmt"

func shiftKerja(arr []string) string {
	var pagi []string
	var siang []string
	var malam []string

	for i := 0; i < len(arr); i++ {
		if i%3 == 0 {
			pagi = append(pagi, arr[i])
		} else if i%3 == 1 {
			siang = append(siang, arr[i])
		} else {
			malam = append(malam, arr[i])
		}
	}

	fmt.Println("Shift Pagi :", len(pagi), pagi)
	fmt.Println("Shift Siang :", len(siang), siang)
	fmt.Println("Shift Malam :", len(malam), malam)

	return ""

}

func main() {
	arr := []string{
		"Upik",
		"Romi",
		"Jalal",
		"Sabo",
		"Renata",
		"Syafiq",
		"Bryan",
		"Monica",
		"Linda",
		"Syifa",
		"Reina",
		"Alfred",
		"Wahid",
		"Fatih",
		"Usop",
		"Shinta",
		"Patrick",
		"Sandy"}

	fmt.Println(shiftKerja(arr))
}
