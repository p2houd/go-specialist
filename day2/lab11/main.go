package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	fmt.Printf("Type a %T size %d bytes in memory\n", a, unsafe.Sizeof(a))
	fmt.Printf("Type b %T size %d bytes in memory\n", b, unsafe.Sizeof(b))

	var usualInt int = 123_567
	var c64 int64 = 16_213_897
	fmt.Println(c64 + int64(usualInt))

	fname := "Abc"
	lname := "Dfe"
	fl := fname + " " + lname
	fmt.Println("FL: ", fl)
	fmt.Println("Length of string", len(fl))
	fl = fl + "Ы"
	fmt.Println("FL: ", fl)
	fmt.Println("Length of string plus utf8 rune:", len(fl))
	fmt.Println("Length of string in runes:", utf8.RuneCountInString(fl))

	totalString, subString := "ABCD", "AB"
	fmt.Println(strings.Contains(totalString, subString))

	// тут исключение из правила. одна из переменных новая а одна нет. но используем :=
	totalString, newSubString := "ЫABCD", "Ы"
	fmt.Println(strings.Contains(totalString, newSubString))

	var sampleRune rune = 'a'
	var anotherRune = 'Q'
	var thirdRune = 77
	fmt.Println(sampleRune, anotherRune, thirdRune)
	fmt.Printf("%c %c %c\n", sampleRune, anotherRune, thirdRune)
	fmt.Printf("%d %d %d\n", sampleRune, anotherRune, thirdRune)

	var aByte byte = 255
	fmt.Println("byte", aByte)

	// in python bug
	// >>> 0.1 + 0.1 + 0.1 == 0.3
	// False
	// >>> 0.1 + 0.1 + 0.1
	// 0.30000000000000004
	var floatBug float64
	floatBug = 0.1 + 0.1 + 0.1
	println(floatBug == 0.3)

	floatBug = 0.011 + 0.089 + 0.1 + 0.3
	println(floatBug == 0.5)

	numA := 0.2
	numB := 0.4
	result := numA + numB
	fmt.Println("%f \n", result)
	fmt.Println(result == 0.6)

	result2 := 0.2 + 0.4
	fmt.Printf("%T %f\n", result2, result2)
	fmt.Println(result2 == 0.6)
}
