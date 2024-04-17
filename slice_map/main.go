package main

import "fmt"

func main() {
	students := []map[string]string{
		{"Nama": "Riko", "Score": "A"},
		{"Nama": "Rika", "Score": "B"},
		{"Nama": "Raka", "Score": "D"},
	}
	for _, student := range students { // supaya berurutan dicetaknya
		fmt.Println(student["Nama"], " - ", student["Score"])
	}
}
