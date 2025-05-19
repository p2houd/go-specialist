package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	var number2 float64
	fmt.Scan(&number)

	fmt.Printf("Первая %d, Вторая %d. Третья %d\n", number/100, (number-number/100*100)/10, number%(number-(number/100*100)/10))

	fmt.Printf("Первая %d, Вторая %d. Третья %d\n", number/100, number/10%10, number%10)

	number2 = math.Pow(1, 2)

	fmt.Printf("abc:%.3f\n", number2)
}
