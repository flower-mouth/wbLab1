/*К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}*/

package task15

import (
	"fmt"
	"strconv"
)

var justString string

func createHugeString(val int) string {
	var str string
	for i := 0; i < val; i++ {
		str += strconv.Itoa(i%10) + "_"
	}
	return str
}

// строка - это массив байтов, следовательно срез будет идти не посимвольно, а побайтно
func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Printf("Huge string:%v\n", v)
	justString = string([]rune(v)[:100])
}

func Task15() {
	someFunc()
	fmt.Printf("Just string:%v", justString)
}
