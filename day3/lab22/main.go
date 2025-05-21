package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 100
	array[1] = 200
	array[2] = 300

	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2]
	fmt.Println("Slice[0:2]: ", startSlice)

	thirdSlice := startArr[1:4]
	for i, _ := range thirdSlice {
		thirdSlice[i]++
	}
	// изменил и оригинальный массив и новый - это слайсы
	fmt.Println(thirdSlice)
	fmt.Println(startArr)

	originArr := []int{1, 2, 3, 34, 4, 3, 5, 43, 534, 5}
	// один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]
	fmt.Println("Before changing. Arr", originArr, "fslice", fSlice, "sslice", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After changing. Arr", originArr, "fslice", fSlice, "sslice", sSlice)

	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "cap:", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice), "cap:", cap(wordsSlice))
	// при этом cap будет больше, берет с запасом x2. capacity
	// показывает сколько элементов можно хранить в слайсе не выделяя дополнительной памяти

	// по степеням двойки еще расширяет 128 .. 256 и тд

	// после выделения памяти под новый массив со старого будут перенесены в новый
	numArr := []int{1, 2}
	numSlice := numArr[:]
	numSlice = append(numSlice, 3) // в этот момент он больше не ссылается на numArr
	numSlice[0] = 0
	fmt.Println(numArr)
	fmt.Println(numSlice)
	// copy()
	// make более детально позволяет создавать срезы

	sl := make([]int, 10, 15)
	fmt.Println(sl)

	myWords := []string{"one", "two", "three"}
	fmt.Printf("myWords:%s,%p\n", myWords, myWords)
	anotherWords := []string{"four", "five", "six"}
	myWords = append(myWords, anotherWords...)
	fmt.Printf("myWords:%s,%p\n", myWords, myWords)

	// вычисление емкости строки бессмысленно тк строка базовый тип
	fmt.Println(cap([]rune("Вася")))

	name := []rune("John")
	name[0] = 'B'
	fmt.Println("Rune aftrert change", string(name))

	myDecimalByteSlice := []byte{101, 101, 102, 103}
	myDecimalString := string(myDecimalByteSlice)
	fmt.Println("myDecimalStr:", myDecimalString)

	// строки неизменяемые а слайсы изменяемые

}
