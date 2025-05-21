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
			continue
		}
		i++
		space++
		if utf8.RuneCountInString(maxWord) > maxRunes {
			maxRunes = utf8.RuneCountInString(maxWord)
		}
		if space >= 5 {
			break
		}
		// вопрос что делать с последним словом и как прерывать
		// мы идем по каждому элементу строки и если это не пробел то конструируем новое слово
		// если это пробел то мы считаем пробелы
	}
	fmt.Printf("%s, %s\n", maxWord, maxRunes)
}
