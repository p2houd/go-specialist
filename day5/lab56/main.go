package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	fmt.Println(a == b)
	i := []int{5, 6, 7}
	i[0] = 19
	fmt.Println(i[0])
	// comp error mismatched type
	// c := 5
	// d := 5.1
	// fmt.Println(c + d)
	j := 5
	defer hello(j)
	j += 10
}
