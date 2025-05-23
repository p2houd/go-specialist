/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

Реализовать конструктор и несколько методов у типа "Накладная"

Пример:
invoice = NewInvoice()

или

order = NewOrder()

*/

package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type Pattern struct {
	name  string
	value string
}

type LengthRequirement struct {
	min int
	max int
}

type Validator struct {
	field string
	pat   Pattern
	req   LengthRequirement
}

func CheckField(s string, pattern Pattern, req LengthRequirement) (bool, error) {
	sLength := utf8.RuneCountInString(s)
	if !(sLength >= req.min && sLength <= req.max) {
		e := fmt.Sprintf("Field %s does not meet length requirements (%v to %v)", s, req.min, req.max)
		log.Println(e)
		return false, errors.New(e)
	}
	var hit int
	for _, v := range s {
		if strings.Contains(pattern.value, string(v)) {
			hit++
		}
	}
	if hit != sLength {
		e := fmt.Sprintf("Field %s does not meet pattern %v", s, pattern.name)
		log.Println(e)
		return false, errors.New(e)
	}
	return true, nil
}

type Address struct {
	ind    string
	city   string
	street string
	bld    string
	fl     string
}

type Order struct {
	name     string
	quantity string
	fio      string
	tel      string
	addr     Address
}

func (o Order) Show() {
	fmt.Println(o.name)
}

func (o Order) New(
	name string,
	quantity string,
	fio string,
	tel string,
	ind string,
	city string,
	street string,
	bld string,
	fl string,
) *Order {
	// fields := []string{name, quantity, fio, tel, addr.city}
	fields := map[string]Validator{
		name: Validator{
			field: name,
			req: LengthRequirement{
				min: 1,
				max: 100,
			},
			pat: Pattern{"letters", "abcdefghiklmnopqrstvxyz"},
		},
		quantity: Validator{
			field: quantity,
			req: LengthRequirement{
				min: 1,
				max: 10,
			},
			pat: Pattern{"digits", "123456789"},
		},
		fio: Validator{
			field: fio,
			req: LengthRequirement{
				min: 1,
				max: 100,
			},
			pat: Pattern{"letters", "abcdefghiklmnopqrstvxyz"},
		},
		tel: Validator{
			field: tel,
			req: LengthRequirement{
				min: 1,
				max: 10,
			},
			pat: Pattern{"digits", "0123456789"},
		},
		ind: Validator{
			field: ind,
			req: LengthRequirement{
				min: 1,
				max: 6,
			},
			pat: Pattern{"digits", "0123456789"},
		},
	}
	for _, v := range fields {
		_, e := CheckField(v.field, v.pat, v.req)
		if e != nil {
			return nil
		}
	}
	return &Order{name, quantity, fio, tel, Address{ind, city, street, bld, fl}}
}

type Item interface {
	Show()
	New()
}

// Ваш код

func main() {
	o := new(Order)
	o1 := o.New("cocacola", "dsf", "ivan", "8800200060", "111500", "Moscow", "Vorontsovskaya", "35", "5")
	if o1 != nil {
		o1.Show()
	}
	o2 := o.New("cocacola", "0", "ivan", "8800200060", "111500", "Moscow", "Vorontsovskaya", "35", "5")
	if o2 != nil {
		o2.Show()
	}
}
