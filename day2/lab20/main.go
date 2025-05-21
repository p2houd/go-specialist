package main

import "fmt"

func main() {
	var summa float64
	// var floatArray [...]float64 = {1,2,3}
	floatArray := [...]float64{1, 2, 3}
	for _, value := range floatArray {
		summa += value
	}
	fmt.Println(summa)

	// как это заставить работать?
	// var floatArray [3]float64 = {1,2,3}
	for _, value := range floatArray {
		summa += value
	}
	fmt.Println(summa)

	// многомерные массивы

	words := [2][2]string{
		{"Mark", "Alice"},
		{"Viktor", "Anna"},
	}
	fmt.Println(words)

	for _, innerAr := range words {
		for _, word := range innerAr {
			fmt.Printf("Word: %s\n", word)
		}
	}
}
