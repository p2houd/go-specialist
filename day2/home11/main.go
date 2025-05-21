package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введи 5 чисел")
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	words := strings.Fields(s)
	for i, v := range words {
		fmt.Println(v, i)
	}
}
