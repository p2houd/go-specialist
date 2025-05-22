package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "proces suces"
}

func main() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("Recioved v;laied", v)
			return
		default:
			fmt.Println("no v alie recievd")
		}
	}
}

// ждем пока не получили данные
// добавление дедлока стразхует нас в ходе выполнения и берет работу на себя