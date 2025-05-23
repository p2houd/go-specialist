package main

import (
	"fmt"
)


// тут он выдаст 0 что нет значения
func main() {
	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
}
