package main

import "fmt"

const minus = 1

func main() {
	var (
		age  int
		name string
	)
	fmt.Println("Enter name:")
	fmt.Scan(&name)
	fmt.Println("Enter age:")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s.\nYour age is %d.\nWhat do you want from me?\n", name, age-minus)
}
