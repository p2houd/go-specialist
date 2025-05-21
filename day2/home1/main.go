package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	// var prev int
	// var ascCount, descCount, constCount int
	fmt.Println("Введи 5 чисел")
	fmt.Scan(&a)
	arr := strings.Split(a, ",")
	var arr2 []int
	for i, v := range arr {
		m, _ := strconv.Atoi(v)
		fmt.Println(m, i)
	}
	fmt.Println(arr2)
	// fmt.Println(prev)
	// for i := 1; i < 5; i++ {
	// 	d := int(a[i] - '0')
	// 	if d > prev {
	// 		ascCount++
	// 	} else if d < prev {
	// 		descCount++
	// 	} else if d == prev {
	// 		constCount++
	// 	}
	// 	prev = d
	// }
	// fmt.Println(ascCount)
	// fmt.Println(descCount)
	// fmt.Println(constCount)
}
