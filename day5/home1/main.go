/*
Задача №2

Вход:
Пользователь должен ввести "правильный пароль", состоящий из:
цифр, букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что пароль правильный и он принят.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8

Реализацию оформить через функцию:
1. checkPassword(pass string) (bool, errors <- на усмотрение)
*/
package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Ваш код

type Pattern struct {
	name  string
	value string
}

type PassLengthRequirement struct {
	min int
	max int
}

func checkPassword(pass string, pattern Pattern) (bool, error) {
	passLength := utf8.RuneCountInString(pass)
	req := PassLengthRequirement{
		min: 8,
		max: 15,
	}
	if !(passLength >= req.min && passLength <= req.max) {
		return false, errors.New(fmt.Sprintf("Password does not meet length requirements (%v to %v)", req.min, req.max))
	}
	var miss int
	for _, v := range pass {
		if !strings.Contains(pattern.value, string(v)) {
			miss++
		}
	}
	if miss != passLength {
		return true, nil
	}
	return false, errors.New(fmt.Sprintf("Pass should have at least one symbol for pattern %v", pattern.name))
}

func checkPatterns(pass string) (bool, error) {
	var patterns []Pattern
	var r bool
	var e error
	patterns = []Pattern{
		Pattern{"digits", "0123456789"},
		Pattern{"lowercase", "abcdefghiklmnopqrstvxyz"},
		Pattern{"uppercase", "ABCDEFGHIKLMNOPQRSTVXYZ"},
		Pattern{"special", "_!@#$%^&"},
	}
	for _, v := range patterns {
		r, e = checkPassword(pass, v)
		if e != nil {
			fmt.Println(r, e)
			break
		}
	}
	return r, e
}
func main() {
	pass := "a$1dsfsFbc"
	checkPatterns(pass)
}
