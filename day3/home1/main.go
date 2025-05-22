/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через другую строку.

Задача №1.1
Реализовать и функцию зашифровки

codeToString(code) -> "???????'

stringToCode("hello") -> "??????"
*/
package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode/utf8"
)

// Ваш код

func codeToString(c string) string {
	// для сдвига Ascii
	var shiftA int64 = 97
	if utf8.RuneCountInString(c)%2 != 0 {
		log.Fatal("String should be even")
	}
	c += " "
	var pair []rune
	var s string
	for _, v := range c {
		if len(pair) < 2 {
			pair = append(pair[:], v)
			continue
		}
		p, _ := strconv.ParseInt(string(pair[:]), 10, 8)
		if p == 26 {
			s += " "
		} else {
			s += string(byte(p + shiftA))
		}

		pair = nil
		pair = append(pair[:], v)
	}
	return s
}

func stringToCode(s string) string {
	// для сдвига Ascii
	var shiftA int64 = 97
	var c string
	for _, v := range s {
		if v != ' ' {
			c += fmt.Sprintf("%02d", v-rune(shiftA))
		} else {
			c += "26"
		}
	}
	return c
}

func main() {
	s := codeToString("220411112603141304")
	// fmt.Printf("%#x\n", s)
	fmt.Printf("%s\n", s)
	s2 := stringToCode("well done")
	// fmt.Printf("%#x\n", s2)
	fmt.Printf("%s\n", s2)
}
