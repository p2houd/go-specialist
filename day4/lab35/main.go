package main

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

func (e Employee) SetName(newName string) {
	e.name = newName
}

func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

func main() {
	e := Employee{
		name:     "A",
		position: "B",
		salary:   99999,
		currency: "D",
	}

	e.SetName("Alex")
	e.SetSalary(3400)

	// p := printers.ResourcePrinter
	// printers.ResourcePrinter.PrintObj(e, os.Stdout)

	// Аналогично setsalary
	(&e).SetSalary(4500)

	// это с точки зрения безопасности может лучше и скопировать чтобы его не перезаписать
	// можно использовать методы с поинт ресивером в случае когда
	// 1) изменения в работе метода должны быть видны на вызывающей стороне
	// когда достаточно увесистый то есть копировать дорого
	// с ними может работать любой вид экземпляра
}
