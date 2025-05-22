package main

import "fmt"

func Describe(pretendent interface{}) {
	fmt.Printf("Type=%T,value=%v\n", pretendent, pretendent)
}

type Describer interface {
	Describe()
}

// мы создали два типа, у них метода create
// но хотим сделать поведение по умолчанию что типа если нет такого метода то вызывай метод из типа по умолчанию
// так вот тогда делаем через несколько интерфейсов и проверяем принадложность к тому или другому

// func (d Describer) Sum() {
// 	fmt.Println("Summing")
// }

type Student struct {
	name string
}

func (st Student) Describe() {
	fmt.Printf("On student")
}

type FigBuilder interface {
	// набор сигнатур методов которые нужно реализовать в другом типе в структуре претенденте
	// во первых у претенденте может быть метод Area() возвращающий int
	Area() int
	// во вторых метод перимере возвращаютщий int
	Perimeter() int
	// PrintObj() error
}

func Area(i FigBuilder) int {
	return 1
}

type Rectangle struct {
	width, length int
}

func (r Rectangle) Area() int {
	return r.length + r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

type Circle struct {
	radius int
}

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func TypeFinder(i interface{}) {
	// присваиваем переменной v значения лежащие под предполагаемым интерфайсом
	switch v := i.(type) {
	case string:
		fmt.Println("eto string", v)
	case int:
		fmt.Println("eto ne string", v)
	case Student:
		fmt.Println("I am student")
		v.Describe()
	case Describer:
		fmt.Println("I am describler")
		v.Describe()
	default:
		fmt.Println("I dont understand u")
	}

}

func main() {
	std := Student{"a"}
	TypeFinder(std)

}
