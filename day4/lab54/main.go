package main

import (
	"fmt"
	"sync"
)

// разрешим состояние гонок при помощи добавления состиоянию мотюкс
// в каналах там все равно используются мютексы

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value for x", x)
}
