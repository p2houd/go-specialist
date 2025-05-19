package main

import (
	"fmt"
	"math"
)

func main() {
	var number2 float64 = math.Pow(1, 2)
	var number3 int = 16

	fmt.Printf("abc:%.3f\n", number2)
	fmt.Printf("abc:%b\n", number3)
	fmt.Printf("abc:%#x\n", number3)
}
