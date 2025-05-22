package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

type Tutor struct {
	name     string
	position string
	salary   int
	currency string
}

// методы те же фунеции но привязанные к определенным структурам

// шаблон func(s Struct) MetodName(params type) output type() {body}
// receiver получатель метода
func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
}

func (t Tutor) DisplayInfo() {
	fmt.Println("Name:", t.name)
	fmt.Println("Position:", t.position)
}

// multiple receivers так нельзя ошибка
// func (t Tutor, e Employee) DisplayInfo() {
// 	fmt.Println("Name:", t.name)
// 	fmt.Println("Position:", t.position)
// }

// запрещено создание функций с одинаковыми названиями
// а методы можно создавать с одиаковыми названиями

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// функция
func Perimeter(c Circle) float64 {
	return 2 * c.radius * math.Pi
}

// метод
func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return float64(r.length) * float64(r.width)
}

func main() {
	e := Employee{
		name:     "A",
		position: "B",
		salary:   99999,
		currency: "D",
	}
	e.DisplayInfo()
	c := Circle{
		radius: 10,
	}
	fmt.Println(Perimeter(c))
	fmt.Println(c.radius)
	fmt.Println(c.Perimeter())
	r := Rectangle{
		width:  10,
		length: 10,
	}
	fmt.Println(r.Perimeter())
}
