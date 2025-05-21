package main

import (
	"fmt"
)

func main() {
	var a int
	var s int
	fmt.Printf("Какие у вас оценки?\n")
	for m := 0; m <= 3; m++ {
		fmt.Scan(&a)
		s += a
	}
	fmt.Printf("Ваши баллы: %d\n", s)
}
