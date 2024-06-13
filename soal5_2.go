package main

import "fmt"

func solution(noAkun []int, nominal []int, isDebit []bool) string {
	totalDebit := 0
	totalCredit := 0
	result := ""

	for i := 0; i < len(noAkun); i++ {
		debit := 0
		credit := 0
		if isDebit[i] {
			debit = nominal[i]
			totalDebit += debit
		} else {
			credit = nominal[i]
			totalCredit += credit
		}

		result += fmt.Sprintf("%d - %d - %d\n", noAkun[i], debit, credit)
	}
	result += fmt.Sprintf("Total = %d %d\n", totalDebit, totalCredit)

	if totalDebit == totalCredit {
		result += "(Balance)"
	} else {
		result += "(Not Balance)"
	}
	fmt.Println(result)
	return result
}

func main() {
	solution([]int{111, 211, 201},
		[]int{200000, 50000, 150000}, []bool{true, false, false})
	solution([]int{111, 201},
		[]int{100000, 120000}, []bool{true, false})
}
