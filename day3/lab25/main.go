package main

import (
	"fmt"
	"strings"
)

func main() {
	// Map набор ключ значение
	// Инициализация пустой мапы

	mapper := make(map[string]int)

	fmt.Println()

	fmt.Println("Empty map:", mapper)

	mapper["Petr"] = 23
	mapper["Elena"] = 12

	fmt.Println("After addring paras in mapa", mapper)

	// ключом в мапе может быть только сравнимный тип (==, !=)
	// мапа не упорядочена
	fmt.Println(mapper["Petr"])
	fmt.Println(mapper["Pavel"])

	if value, ok := mapper["Petr"]; ok {
		fmt.Printf("Denis exist his name is %d\n", value)
	}

	fmt.Println(strings.Repeat("=", 15))
	for key, value := range mapper {
		fmt.Printf("key: %s and value %v\n", key, value)
	}

	// как удалять пары
	fmt.Println("Befor deliting", mapper)
	delete(mapper, "Petr")
	fmt.Println("After deliting", mapper)

	// delete дорогая операция

	// слайсы ссылочный тип опэтому их можно сравнивать только с nil на равенство

	// if []int{1,2,3} == []int{1,2,3}{
	// 	fmt.Println()
	// }

	// и мапы тоже
	// но можно взять поинтер и сравнить

	if &[]int{1, 2, 3} == &[]int{1, 2, 3} {
		fmt.Println()
	}
}
