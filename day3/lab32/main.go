package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Befor", arr)
	mutation(&arr)
	fmt.Println("After", arr)
	// вот так лучше использовать
	sliceMutation(arr[:])
	fmt.Println("After", arr)
}

func mutation(arr *[3]int) {
	// (*arr)[1] = 900
	// (*arr)[2] = 1000
	// НО можно писать и так потому что GO и сам разименует указатель на массив

	arr[1] = 900
	arr[2] = 1000
}

// используйте слайсы - это идеоматично с точки зрения golang

func sliceMutation(slice []int) {
	slice[1] = 900
	slice[1] = 1000
}
