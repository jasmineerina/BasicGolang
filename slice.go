package main

import "fmt"

func main() {
	names := [...]string{"Yoo", "Jimin", "Ning", "Yizhuo", "Aeri", "Uchinaga", "Kim", "Minjeong"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:2]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)
}
