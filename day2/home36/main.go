package main

import "fmt"

func main() {
	var a string
	var s int
	fmt.Scan(&a)
	for i := 0; i < len(a); i++ {
		s += int(a[i] - '0')
	}
	fmt.Printf("Сумма цифр вашего числа: %d\n", s)
}
