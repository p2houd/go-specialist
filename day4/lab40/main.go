package main

import (
	"fmt"
	"time"
)

func newGoRoutine(a int) {
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Println("Hey i am gorutine", a)
}

func main() {
	go newGoRoutine(1)
	go newGoRoutine(2)
	fmt.Println("Main gorutine")
	time.Sleep(3 * time.Second)
}
