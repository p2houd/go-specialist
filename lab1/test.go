package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	x := 5
	y := 7
	sum := add(x, y)
	fmt.Printf("Сумма %d и %d равна %d\n", x, y, sum)
}
