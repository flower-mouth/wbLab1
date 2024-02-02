/*Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false*/

package task26

import (
	"fmt"
	"strings"
)

func check(str string) bool {
	hashMap := make(map[rune]int)
	for _, value := range str {
		hashMap[value]++
		if hashMap[value] > 1 {
			return false
		}
	}
	return true
}

func Task26() {
	input := "abCdefAaf"
	input = strings.ToLower(input)
	fmt.Printf("%v - %v", input, check(input))
}
