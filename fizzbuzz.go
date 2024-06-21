package main

import "fmt"

func cariBilangan(angka []int) string {
	var FizzBuzz []int
	var Fizz []int
	var Buzz []int

	for i := 0; i < len(angka); i++ {
		if angka[i]%6 == 0 && angka[i]%7 == 0 {
			FizzBuzz = append(FizzBuzz, angka[i])
		} else if angka[i]%6 == 0 {
			Fizz = append(Fizz, angka[i])
		} else if angka[i]%7 == 0 {
			Buzz = append(Buzz, angka[i])
		}
	}

	fmt.Println("Fizz", Fizz)
	fmt.Println("Buzz", Buzz)
	fmt.Println("FizzBuzz", FizzBuzz)

	return ""
}

func main() {
	angka := []int{30, 33, 34, 35, 36, 38, 40, 42, 43, 45, 49}

	fmt.Println(cariBilangan(angka)) // fizz : 30, 36 || buzz :  35, 49 || fizzbuzz : 42
}
