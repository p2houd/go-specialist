package main

import (
	"fmt"
)

type University struct {
	city string
	name string
}

func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

type Professor struct {
	name string
	age  int
	University
}

func main() {
	prof := Professor{
		name: "Boris Bob",
		age:  42,
		University: University{
			city: "Eng",
			name: "BST",
		},
	}
	prof.FullInfoUniversity()
}
