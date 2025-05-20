package main

import "fmt"

func main() {
	var (
		deposit float64
		// years       int64
		// percent     int64
		// result      float64
		depositHint string = "Размер вклада от 100 до 1 000 000"
		// yearsHint   string = "Количество лет от 1 до 100"
		// percentHint string = "Процент по вкладу от 1 до 20"
		error string = "Неправильно!"
	)

	fmt.Println(depositHint)
	fmt.Scan(&deposit)
	if deposit < 100 || deposit > 1000000 {
		fmt.Println(error, depositHint)
		return
	}
	fmt.Println("Все нормально. Депозит", deposit)
}
