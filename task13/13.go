//Поменять местами два числа без создания временной переменной.

package task13

import "fmt"

func Task13() {
	a := 0
	b := 1
	a, b = b, a
	fmt.Printf("a=%v, b=%v", a, b)
}
