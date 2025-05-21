package main

import (
	"fmt"
)

func main() {
	// массивы

	var arr [5]int // при инициализации массива важно передать информацию сколько элементов в нем

	fmt.Println("This is array", arr)

	// присваивание значения элементу массива
	arr[0] = 10
	arr[1] = 20
	arr[3] = -30

	fmt.Println("Values", arr)
	// hex
	fmt.Printf("%#x\n", arr)

	newArr := [5]int{1, 3, 4}
	fmt.Printf("%#x\n", newArr)

	// массивы не зависят друг от друга, создается новая копия

	arrVal := [...]int{10, 20, 30, 40}
	fmt.Printf("%#x\n", arrVal)
	fmt.Printf("%d\n", len(arrVal))

	second := arrVal
	second[1] = 1000
	fmt.Println(arrVal)
	fmt.Println(second)

	// массив и его размер  это две составляющие одного типа, то есть массивы разных размеров это разные массивы

	// итерирование по массиву
	floatArray := [...]float64{12.5, 10.1, 12.1}
	var sum float64
	for i := 0; i < len(floatArray); i++ {
		fmt.Printf("%d element of array is %.2f\n", i+1, floatArray[i])
		sum += floatArray[i]
	}
	fmt.Println(sum)
}
