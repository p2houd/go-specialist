package main

import "fmt"

type rect struct {
	len, wid int
}

func (r rect) area() {
	fmt.Println(r.len * r.wid)
}

func main() {
	r := &rect{
		len: 10,
		wid: 10,
	}
	r.area()
	fmt.Println(string(65))
}
