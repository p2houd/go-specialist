package main

import (
	"fmt"
)

func main() {
	var variable int = 30
	var pointer *int = &variable

	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Printf("Value of pointer: %v\n", pointer)

	var zeroPointer *int

	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)

	// создание указателей на нулевые значения
	var zeroVar int
	var zeroPoint *int = &zeroVar
	fmt.Println(zeroPoint)
	zeroNewPoint := new(int)
	fmt.Printf("value of zeronewpoint: %v and adres %v\n", *zeroNewPoint, zeroNewPoint)
	*zeroNewPoint++
	fmt.Printf("value of zeronewpoint: %v and adres %v\n", *zeroNewPoint, zeroNewPoint)

}
