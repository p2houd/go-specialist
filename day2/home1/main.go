package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var prev int
	var ascCount, descCount, constCount int
	fmt.Println("Введи 5 чисел")
	fmt.Scan(&a)
	arr := strings.Split(a, ",")
	prev = int(arr[0] - '0')
	fmt.Println(prev)
	for i := 1; i < 5; i++ {
		d := int(a[i] - '0')
		if d > prev {
			ascCount++
		} else if d < prev {
			descCount++
		} else if d == prev {
			constCount++
		}
		prev = d
	}
	fmt.Println(ascCount)
	fmt.Println(descCount)
	fmt.Println(constCount)
}
