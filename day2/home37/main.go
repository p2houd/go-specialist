package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string
	var maxRunes int
	var curWord []rune
	var maxWord string
	var space int
	// s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s = "1 22 333 4444 55555 666666"
	for _, v := range s {
		// fmt.Println(i, v)
		if v != ' ' {
			curWord = append(curWord[:], v)
			continue
		}
		// fmt.Println(string(curWord))
		if utf8.RuneCountInString(string(curWord)) > maxRunes {
			maxRunes = utf8.RuneCountInString(string(curWord))
			maxWord = string(curWord)
		}
		// fmt.Println(maxWord)
		if space > 5 {
			break
		}
		curWord = nil
		space++
	}
	fmt.Printf("%v, %v\n", maxWord, maxRunes)
}
