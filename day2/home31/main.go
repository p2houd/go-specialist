package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	var b int
	var s string
	fmt.Scan(&a)
	fmt.Scan(&b)
	if b < a {
		fmt.Println("Нет, b должно быть больше a")
		return
	}
	for v := a; v <= b; v++ {
		s += fmt.Sprintf("%d,", v)
	}
	s, _ = strings.CutSuffix(s, ",")
	fmt.Println(s)
}
