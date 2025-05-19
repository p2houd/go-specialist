package main

import (
	"fmt"
)

func main() {
	var result float64
	var d, e float64 = 10, 4 // тут d и e не целые числа, поэтому и результат не целое

	var check float64 = 20 / 3 // тут используется два целых числа, поэтому результат целое число

	a := 42
	b := 20
	c := a / b
	result = d / e
	result2 := d / 3

	fmt.Printf("type of c %T and value %v\n", c, c)
	fmt.Printf("type of result %T and value %v\n", result, result)
	fmt.Printf("type of result2 %T and value %v\n", result2, result2)
	fmt.Printf("type of check %T and value %v\n", check, check)

	result2++
	fmt.Printf("type of result2 %T and value %v\n", result2, result2)

	result2--
	fmt.Printf("type of result2 %T and value %v\n", result2, result2)

	// возьмем остаток float64 не имеет оператора %, поэтому приводим к int
	newCheck := int(check) % 3
	fmt.Println("newCheck:", newCheck)
}
