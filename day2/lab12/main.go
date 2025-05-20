package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 5; i++ {
		if i > 3 {
			break
		}
		fmt.Println("Value is:", i)
	}
	fmt.Println("Full is finished")

	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Value is:", i)
	}
	fmt.Println("Full is finished")
}
