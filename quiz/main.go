package main

import "fmt"

func main() {
	// hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	hasilScore := 0

	// for in range
	for _, hasil := range scores {
		hasilScore += hasil
	}
	fmt.Println("Total Score :", hasilScore)

	hasilScore /= len(scores)

	fmt.Println("Rata rata :", hasilScore)

}
