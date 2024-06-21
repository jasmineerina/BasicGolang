package main

import "fmt"

// Can Kelly solve enough code challenges to catch up with Sam?
// Sam and Kelly are programming buddies.
// Kelly resolves to practice more as Sam is ahead initially.
// They each solve a number of problems daily.
// Find the minimum number of days for Kelly to have solved more problems than Sam.
// If Kelly cannot surpass return -1.

func minNum(samDaily int, kellyDaily int, diff int) int {
	samSolved := diff + samDaily
	kellySolved := kellyDaily
	result := 1

	if samDaily == kellyDaily {
		result = -1
		return result
	}

	for samSolved > kellySolved {
		samSolved += samDaily
		kellySolved += kellyDaily
		result++
	}

	return result

}

func main() {
	fmt.Println(minNum(3, 5, 5))
	fmt.Println(minNum(3, 5, 1))
}
