package main

import "fmt"

func main() {
	// case с проваливаниями даже ложные кейсы
	// В момент выполнения инструкции fallthrough

	var num int
	fmt.Scan(&num)

	// если первое выполнилось то и остальные тоже выполнились и даже не проверяются
	// 

outer:
	switch {
	case num < 100:
		fmt.Printf("%d less than 100\n", num)
		if num%2 == 0 {
			break outer
		}
		fallthrough
	case num > 200:
		fmt.Printf("%d greater than 200\n", num)
		fallthrough
	default:
		fmt.Println("Default")
	}
}
