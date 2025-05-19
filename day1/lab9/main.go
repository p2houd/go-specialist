package main

import (
	"fmt"
)

func main() {
	if width := 100; width > 100 {
		fmt.Println("windth >= 100")
		return
	}
	fmt.Println("windht < 100")
}
