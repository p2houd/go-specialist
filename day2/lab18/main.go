package main

import "fmt"

func main() {

	var i int
outloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value i is %d and it is even\n", i)
			break outloop
		}
	}
	fmt.Println("The ends")
}
