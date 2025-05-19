package main

import (
	"fmt"
)

func main() {
	// Boolean default false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// Boolean operands
	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	// Классический условный оператор

	var value int
	_, e := fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("chetny")
	} else if value == 5 {
		fmt.Println("pyatorka")
	} else {
		fmt.Println("ne chenty")
	}
	if e != nil {
		fmt.Printf("%s\n", e)
	}

}
