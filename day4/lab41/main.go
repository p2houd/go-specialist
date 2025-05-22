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
	// done <- 1
}

func main() {
	var a chan int
	// var r int
	if a == nil {
		fmt.Println("chanel is nil, letz defin it")
		a = make(chan int)
		fmt.Printf("Typ of a is %T\n", a)
	}
	go newGoRoutine(1, a)
	go newGoRoutine(2, a)
	<-a
	fmt.Println("Main gorutine")
	time.Sleep(3 * time.Second)
}

// тут получем дедлок
// дедлок это когда кто то пишет в канал но из него никогда никто не читает
// дедлок это когда кто то чтает из капнала но никто никогда не пишет
