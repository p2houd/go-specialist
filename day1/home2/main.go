/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

// Ваш код

import "fmt"

func main() {
	var aNumber int
	fmt.Println("Введите число, которое нужно перевернуть:")
	fmt.Scan(&aNumber)
	d1 := aNumber % 10
	d2 := aNumber / 10 % 10
	d3 := aNumber / 100 % 10
	fmt.Printf("Ваше перевернутое число: %d%d%d\n", d1, d2, d3)
}
