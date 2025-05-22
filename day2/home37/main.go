package main

import (
	"fmt"
	"unicode/utf8"
)

func findLongestWord(s string) (string, int) {
	var maxRunes int
	var curWord []rune
	var maxWord string
	var space int
	s += " "
	for _, v := range s {
		// fmt.Println(i, v)
		if v != ' ' {
			curWord = append(curWord[:], v)
			continue
		}
		// fmt.Println(maxWord)
		if space > 4 {
			break
		}
		// fmt.Println(string(curWord))
		if utf8.RuneCountInString(string(curWord)) > maxRunes {
			maxRunes = utf8.RuneCountInString(string(curWord))
			maxWord = string(curWord)
		}
		curWord = nil
		space++
	}
	return maxWord, maxRunes
}

func main() {
	// s, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s := "1 22 333 4444 55555 66666666 sd f"
	maxWord, maxRunes := findLongestWord(s)
	fmt.Printf("%v, %v\n", maxWord, maxRunes)
}
