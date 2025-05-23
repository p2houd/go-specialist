package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:2:5]
	fmt.Println(t[1])
	fmt.Printf("%%\n")
}
