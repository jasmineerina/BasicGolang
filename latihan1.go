package main

import "fmt"

func cekModulus(num string) string {
	isMod := ""

	for i := 0; i < len(num); i++ {
		index := i + 1
		digit := int(num[i] - '0')

		if digit%index == 0 {
			isMod = "true"
		} else {
			isMod = "false"
			break
		}
	}
	return isMod
}

func main() {
	fmt.Println(cekModulus("14"))
	fmt.Println(cekModulus("236"))

}
