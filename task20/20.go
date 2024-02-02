/*Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».*/

package task20

import (
	"fmt"
	"strings"
)

func Task20() {
	str := "snow dog sun"
	strArray := strings.Fields(str)
	fmt.Println(strArray)

	for left, right := 0, len(strArray)-1; left < right; left, right = left+1, right-1 {
		strArray[left], strArray[right] = strArray[right], strArray[left]
	}

	fmt.Printf("Input string: %s\nResult string: %s", str, strings.Join(strArray, " "))
}
