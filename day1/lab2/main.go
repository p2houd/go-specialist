package main

import "fmt"

func main() {
	var age int
	fmt.Println("My age is:", age)
	fmt.Printf("His type is %T\n", age)

	const price, tax float32 = 275.00, 27.50 // типизированные константы
	const quantity, inStock = 2, true        // нетипизированные

	fmt.Println("Total:", 2*quantity*(price+tax))
	// почему instock не подсветил как не используемую?

	val, a := 2, 1
	// а тут подсвечивает, когда не испольщует

	fmt.Println("Results:", val, a)

	var (
		var1 int
		var2 string
	)

	res, err := fmt.Scan(&var1)
	fmt.Scan(&var2)

	var (
		var3 string
	)

	fmt.Scan(&var3)

	fmt.Println(var1, var2, var3)
	fmt.Println(res, err)
	// в var1 записался 0 когда я подал строку в stdin

}
