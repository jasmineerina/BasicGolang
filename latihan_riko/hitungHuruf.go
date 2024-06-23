package main

import "fmt"

func cariHuruf(huruf string) string {
	var a []string
	var b []string
	var c []string

	for i := 0; i < len(huruf); i++ {
		if string(huruf[i]) == string("a") {
			a = append(a, string(huruf[i]))
		} else if string(huruf[i]) == string("b") {
			b = append(b, string(huruf[i]))
		} else {
			c = append(c, string(huruf[i]))
		}
	}

	fmt.Println("Huruf A:", len(a), a)
	fmt.Println("Huruf B:", len(b), b)
	fmt.Println("Huruf C:", len(c), c)

	return "------------"
}

func main() {
	fmt.Println(cariHuruf("aaabbcccaaaac"))
}
