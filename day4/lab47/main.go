package main

import (
	"fmt"
	"time"
)

func generator(ch chan int) {
	for i := 0; i < 25; i++ {
		ch <- i
	}
	close(ch)

}

func main() {
	ch := make(chan int)
	go generator(ch)

	val, ok := <-ch
	fmt.Println(val, ok)
	time.Sleep(time.Second * 4)
	// вот тут мы если закрыли доступ к каналу то реяем доступ к данным и уже не можем его прочитать
	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)
	val, ok = <-ch
	fmt.Println(val, ok)
}
