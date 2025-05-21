package main

import (
	"fmt"
)

func main() {
	/// anon function
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println(anon)
	fmt.Println(anon(1, 2))
	fmt.Println(bigFunction(1, 2))
	fmt.Println(calcAndReturnValidFunction("addition")(1, 2))
	fmt.Println(calcAndReturnValidFunction("addn")(1, 2))
}

func bigFunction(a, b int) int {
	return func(a, b int) int { return a + b + 100 }(a, b)
}

// явные функции в go в принципе любая это экземпляр 1 уровня, поэтому ее можно присваивать переменной

func calcAndReturnValidFunction(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	}
	return func(a, b int) int { return 0 }
}
