package main

import "fmt"

// вложенные структуры (вложение структур)
// это использование одной структуры как тип поля в другой структуре

type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// встроенные структуры мы добавляем поля одной структуры к другой без миени

type Professor struct {
	firstN     string
	lastN      string
	age        int
	greatWorks string
	// papers map[string]string - добавленгие этого поля делате структуры несравнимео
	University
}

func main() {
	student := Student{
		firstName: "Feda",
		lastName:  "Petrov",
		university: University{
			age:       879,
			yearBased: 1899,
			infoShort: "Universtu",
			infoLong:  "has lot of students",
		},
	}
	fmt.Println(student)
	fmt.Println(student.university.age)
	prof := Professor{
		age: 111,
		University: University{
			age: 123,
		},
	}
	prof2 := Professor{
		age: 111,
		University: University{
			age: 123,
		},
	}
	fmt.Println(prof)
	fmt.Println(prof.University.age)
	fmt.Println(prof == prof2)
	// можно сравнивать, но если хотя бы одно из полей структуры не сравнимо то и вся структура тоже
}

// методы с маленькой буквы они приватные. с полшой они публиычные
