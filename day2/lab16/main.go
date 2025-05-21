package main

import (
	"fmt"
)

func main() {
	var price int
	fmt.Scan(&price)

	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Case two")
	}

	// case с множеством вариантов
	var ageGroup string = "d" // возрастные группы
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("First group")
	case "d", "e":
		fmt.Println("Second group")
	default:
		fmt.Println("Ti ne prav")
	}

	var age int
	fmt.Scan(&age)
	switch {
	case age <= 18:
		fmt.Println("Tu yong")
	case age > 18:
		fmt.Println("Tu old")
	default:
		fmt.Println("Kto ty?")
	}

}
