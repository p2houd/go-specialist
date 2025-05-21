package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	if a > 10 || a < 1 {
		fmt.Println("Нет, число должно быть от 1 до 10")
		return
	}
	for m := 1; m <= 10; m++ {
		fmt.Printf("%dx%d=%d\n", a, m, a*m)
	}
}
