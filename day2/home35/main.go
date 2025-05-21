package main

import (
	"fmt"
)

func main() {
	var a int
	var s int
	for {
		fmt.Scan(&a)
		if a == -1 {
			break
		}
		s += a
	}
	fmt.Printf("Сумма чисел: %d\n", s)
}
