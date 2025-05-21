package main

import (
	"fmt"
	"strings"
)

func main() {

	var in, sum int
	var s string

	fmt.Scan(&in)

	for {
		if in == 0 {
			break
		}
		var d int = in % 10
		sum += d
		in /= 10
		s += fmt.Sprintf("%d + ", d)
	}

	fmt.Printf("Сумма числа: %v\n", sum)
	s, _ = strings.CutSuffix(s, " + ")
	fmt.Println(s, "=", sum)

	// strings.Join
	a := strings.Join([]string{"a", "b"}, "\t")
	fmt.Println(a)

	// strings.Builder
	var b strings.Builder
	b.WriteRune('a')
	fmt.Println(b.String())
}
