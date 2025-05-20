package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Там треугольник?")

outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j:%d\n", i, j, i+j)
			if i == j {
				break outer
			}
		}
	}
	fmt.Println("After outer stop")

	var loopVar int = 0
	// while (loopVar < 10) {
	// 	loopVar++
	// }

	for loopVar := 0; loopVar < 10; {
		fmt.Println("In while like loop", loopVar)
		loopVar++
	}

	// из-за порядка обхода он ее не трогал
	fmt.Println(loopVar)

	for loopVar < 10 {
		fmt.Println("In while loke loop", loopVar)
		loopVar++
	}

	fmt.Println(loopVar)

	for loopVar, second := 5, 12; loopVar < 34 && second < 34; {
		fmt.Println(loopVar, second)
		loopVar++
		second++
	}

}
