package main

import (
	"fmt"
)

func hello() []string {
	return nil
}

type person struct {
	name string
}

func main() {
	h := hello
	fmt.Println(h == nil)

	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}
