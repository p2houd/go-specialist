package main

import "fmt"

// структура заименованный набор состиояний
// структура исходя из своего названия отвечает на вопрос - из чего я должна состоять
// чтобы считаться тем типом который декларируется структурой

// структура описывает паттерн состояния
// интерфейс - явно декларированный набор сигнатур поведения чаще в виде набора методов, удовлетворив который он может сичтаться типом который декларирует интерфейс
// отвечает на вопрос что я должен уметь делать чтобы считаться тем типом

type Rectangle struct {
	width, length int
}

func (r Rectangle) Area() int {
	return r.length + r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

type FigBuilder interface {
	// набор сигнатур методов которые нужно реализовать в другом типе в структуре претенденте
	// во первых у претенденте может быть метод Area() возвращающий int
	Area() int
	// во вторых метод перимере возвращаютщий int
	Perimeter() int
	// PrintObj() error
}


type Empty interface{}

func Describer(pretendent interface{}) {
	fmt.Printf("Type=%T,value=%v\n", pretendent, pretendent)
}

type Student struct {
	name string
}

// type assertion func semi generic
func SemiGeneric(pretendent interface{}) {
	value, ok := pretendent.(int) // проверка что претендент инт
	fmt.Println(value, ok)
}

type Circle struct {
	radius int
}

// func (c Circle) PrintObj() error {
// 	p := printers.HumanReadablePrinter()
// 	p.Prin
// }

func (c Circle) Area() int {
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

// декларируем структуры претенденты

func main() {
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	totalSumOfRectangles, totalSumOfCircles := 0, 0

	rectangles := []Rectangle{r1, r2, r3}
	for _, rect := range rectangles {
		totalSumOfRectangles += rect.Area()
	}

	circles := []Circle{c1, c2, c3}
	for _, circ := range circles {
		totalSumOfCircles += circ.Area()
	}
	fmt.Println("Total area is ", totalSumOfCircles, totalSumOfRectangles)

	// теперь воспользуемся интерфейсом который мы создали раньше

	totalSumOfFigures := 0
	figures := []FigBuilder{r1, r2, r3, c1, c2, c3}
	for _, fig := range figures {
		totalSumOfFigures += fig.Area()
	}
	fmt.Println("Total area is ", totalSumOfFigures)

	strSample := "Helo word"
	Describer(strSample)

	intSample := 200
	Describer(intSample)

	// а теперь работа с полу жинеригом
	SemiGeneric(10)
	SemiGeneric("Student feda")
	SemiGeneric(Student{"Fedo"})
	SemiGeneric("Hel world")
}
