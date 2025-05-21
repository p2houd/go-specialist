package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is", result)
		}
	}

	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)

	fmt.Printf("number %d his type %T\n", numInt, numInt)
}
