package main

import "fmt"

func changeParam(value *int) {
	*value += 100
}

func returnPointer() *int {
	var numeric int = 321
	return &numeric // в момент возрата go перемещает данную пременную в кучу
}

func main() {
	sample := 1
	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("Changing value of sample", sample)
	fmt.Println("return pointer", returnPointer())
}
