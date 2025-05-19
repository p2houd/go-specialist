package main

import (
	"fmt"
	"strings"
)

func main() {
	var tsvet string
	_, e := fmt.Scan(&tsvet)

	if strings.Compare(tsvet, "beli") == 0 {
		fmt.Println("haroshi")
	} else if strings.Compare(tsvet, "cherny") == 0 {
		fmt.Println("pyatorka")
	} else {
		fmt.Println("gg")
	}
	if e != nil {
		fmt.Printf("%s\n", e)
	}

}
