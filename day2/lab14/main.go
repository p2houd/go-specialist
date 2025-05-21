package main

import (
	"fmt"
	"strings"
)

func main() {

	var in, sum int
	var s string

	fmt.Scan(&in)

	for in != 0 {
		d := in % 10
		sum += d
		in /= 10
		s += fmt.Sprintf("%d+", d)
	}

	fmt.Printf("Сумма числа: %v\n", sum)
	s, _ = strings.CutSuffix(s, "+")
	fmt.Println(s)
}
