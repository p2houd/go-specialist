package main

import (
	"fmt"
	"time"
)

// каналы соединительные трубы вода тчечер черезр конатл

// по умолчанию они nil
// map, slice, chan создаем через make

func newGoRoutine(a int, done chan int) {
	time.Sleep(time.Duration(a) * time.Second)
	fmt.Println("Hey i am gorutine", a)
	done <- 1
}

func main() {
	var a chan int
	if a == nil {
		fmt.Println("chanel is nil, letz defin it")
		a = make(chan int)
		fmt.Printf("Typ of a is %T\n", a)
	}
	go newGoRoutine(1, a)
	go newGoRoutine(2, a)
	time.Sleep(1 * time.Second)
	r := <-a
	fmt.Println("Main gorutine", r)
}

// а тут мы дождались только первую горутину засчет получения данных из канала и завершили работу
