package main

import "fmt"

// структура заименованный набор полей

type Student struct {
	firstName string
	lastName  string
	age       int
}

// если имеется ряд состояний одинакового типа то можно сделать так
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func PrintStudent(std AnotherStudent) {
	fmt.Println("==========")
	fmt.Println(std.age)
}

func PrintStudent2(std Student) {
	fmt.Println("==========")
	fmt.Println(std.age)
}

type Humane struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func main() {
	student := AnotherStudent{
		firstName:    "A",
		lastName:     "B",
		groupName:    "C",
		age:          10,
		courseNumber: 999,
	}
	fmt.Println(student)
	PrintStudent(student)
	student2 := Student{"Petr", "Ivanov", 19}
	PrintStudent2(student2)
	// student3 := Student{"Petr"} not woprk
	student3 := Student{
		firstName: "Petr",
	}
	PrintStudent2(student3)

	// анонимные структуры (без имени)
	anonStudent := struct {
		age           int
		groupId       int
		professorName string
	}{
		age:           23,
		professorName: "A",
		groupId:       23,
	}
	fmt.Println(anonStudent)
	anonStudent.age++
	fmt.Println(anonStudent)

	studPointer := &Student{
		firstName: "Ig",
		lastName:  "Norm",
		age:       99,
	}
	fmt.Println(studPointer)
	fmt.Println(*studPointer)
	// неявно происходит разименование и обращние к соответствуюещме полдлю
	fmt.Println("Age via .age:", studPointer.age)
	fmt.Println("Age via .age:", (*studPointer).age)

	// создание экзеплюяра с анонимныеми полями но типа каждого типа толко по одному можно
	humane := &Humane{
		firstName: "Huma",
		lastName:  "Mane",
		string:    "9sad",
		int:       99,
		bool:      true,
	}
	fmt.Println(humane)
}
