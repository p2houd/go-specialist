package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string
	var maxRunes int
	var maxWord string
	var space int
	var i int
	fmt.Scan(&s)
	fmt.Println(s)
	for {
		if s[i] != ' ' {
			maxWord += string(s[i])
			i++
			continue
		}
		i = 0
		space++
		if utf8.RuneCountInString(maxWord) > maxRunes {
			maxRunes = utf8.RuneCountInString(maxWord)
		}
		if space >= 5 {
			break
		}
	}
	fmt.Printf("%s, %s\n", maxWord, maxRunes)
}
