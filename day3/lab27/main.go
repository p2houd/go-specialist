package main

import "fmt"

func main() {
	fmt.Println("call function")
	res := add(10, 20)
	fmt.Println("Result of add (10, 20) is", res)
	a, b := namedReturn(19, 19)
	fmt.Println(a)
	fmt.Println(b)
	c, d := funcWithReturn(1, 2)
	fmt.Println(c)
	fmt.Println(d)
	// ошибка ну тогда нет возвращаемого значения
	// check := emptyReturn(123112)
	// fmt.Println(check)
	helloVariadic(10, 20, 23004, 320402340)
	someStrings(1, 1212, "abc", "bef", "eg1231")
	sliceNumbers := []string{"12", "3", "123, 12, 3, 12, 3"}
	someStrings(1, 2, sliceNumbers...)
}

func add(a int, b int) int {
	result := a + b
	return result
}

func namedReturn(width, height int) (perimetr int32, area int) {
	perimetr = int32(2*width + height)
	area = width * height
	return // не нужно указывать возвращаемые значения
}

// при вызове оператора можно сделать несколько return
func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}
	if a == b {
		return a, true
	}
	return 0, false
}

// что если функция вернет если ретурн не указан или пустой

func emptyReturn(a int) {
	fmt.Println("I'm emptyreturn with parametrs", a)
	return
}

// variadic parameteres (континуальные параметры)

func helloVariadic(a ...int) {
	fmt.Printf("value %v\n", a)
}

// смешивание параметров с variadic
// континуальный параметр всегда последний
// один на всю функцию

func someStrings(a, b int, words ...string) {
	fmt.Println("Parametr a:", a, "Paramentr b", b, "Words:", words)

	for _, word := range words {
		if len(word) > a && len(word) < b {
			fmt.Println("nasli ego")
		}
	}
}
