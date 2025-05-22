package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Perimeter() int {
	p := r.length * r.width
	fmt.Println(p)
	return p
}

// для имен конструкторов договорились со следующим названием
// New() если данный конструктор на файл один (в файле присутствует одна структура)
// New<StructName()> - если в файле присутствуют разные структуры

// принято возвращать из конструктора не сам экземпляр а ссылку на него

func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

func main() {
	r := NewRectangle(1, 20)
	r.Perimeter()

	// defer db.Close() поработает в конце работы кода, в любом случае отработат даже если ошибка
}
