package main

import "fmt"

func solution(noAkun []int, nominal []int) string {
	totalDebit := 0
	totalCredit := 0
	result := ""

	for i := 0; i < len(noAkun); i++ {
		debit := nominal[i]
		credit := 0
		if i == 0 {
			debit = nominal[i]
		} else {
			credit = nominal[i]
		}

		totalDebit += debit
		totalCredit += credit
		result += fmt.Sprintf("%d - %d - %d\n", noAkun[i], debit, credit)
	}
	result += fmt.Sprintln("Total =", totalDebit, totalCredit)

	if totalDebit == totalCredit {
		result += "(Balance)"
	} else {
		result += "(Not Balance)"
	}
	fmt.Println(result)
	return result
}

func main() {
	solution([]int{111, 211, 201}, []int{200000, -50000, -150000})
	solution([]int{111, 201}, []int{100000, -120000})
}
