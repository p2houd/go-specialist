package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started goruting", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goriting %v endei\n", i)
	wg.Done()
}

func main() {
	no := 5
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goruting sopted")
}
