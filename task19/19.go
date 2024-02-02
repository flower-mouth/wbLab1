/*Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/

package task19

import "fmt"

func Task19() {
	var str = "главрыба"
	var runeResult []rune
	runeStr := []rune(str)

	for i := len(runeStr) - 1; i >= 0; i-- {
		runeResult = append(runeResult, runeStr[i])
	}

	fmt.Printf("Input string: %s\nResult string: %s", str, string(runeResult))
}
