package main

import "fmt"

//1. Deadlock - ситуация, когда кто-то пишет в канал НО ИЗ НЕГО НИКОГДА НИКТО НИЧЕГО НЕ ПРОЧИТАЕТ,
//  или когда кто-то читает из канала, НО В НЕГО НИКТО НИКОГДА НЕ ЗАПИШЕТ!

// По сути ситуация означает, что для отправляющей стороны отсутствует получатель (с другой стороны никто не ждет данных). И наоборот.

// func main() {
// 	ch := make(chan int)
// 	ch <- 10
// 	// <- ch
// }

func main() {
	ch := make(chan int, 1)
	ch <- 10
	// time.Sleep(time.Second * 4)
	r := <-ch
	ch <- 10
	// там буфер под капотом. пока не прочитаем мы 
	// hchan - то как устроена структура
	close(ch)
	fmt.Println(r)
}
