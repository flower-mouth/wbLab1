/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в
1 или 0.*/

package task8

import (
	"fmt"
	"strconv"
	"time"
)

func viaRunes(input int64, bitNumber int) {
	fmt.Println("\nusing Runes")
	start := time.Now()

	bitValue := 0

	inBinary := strconv.FormatInt(input, 2)
	runes := []rune(inBinary)
	if bitValue == 0 {
		runes[len(runes)-1-bitNumber] = '0'
	} else {
		runes[len(runes)-1-bitNumber] = '1'
	}
	result, _ := strconv.ParseInt(string(runes), 2, 64)
	binaryResult := strconv.FormatInt(result, 2)

	fmt.Printf("Number: %v (0b%v)\nNumber of bit: %v\n", input, inBinary, bitNumber)
	fmt.Printf("Result: %v (0b%v)\n", result, binaryResult)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}

func viaXOR(input int64, bitNumber int) {
	fmt.Println("\nusing XOR")
	start := time.Now()
	var (
		inBinary int64 = 1 << bitNumber
	)
	fmt.Printf("Number: %v (%#b)\nNumber of bit: %v (%#b)\n", input, input, bitNumber, inBinary)

	res := input ^ inBinary
	fmt.Printf("Result: %v (%#b)\n", res, res)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}

func Task8() {
	input := int64(999)
	bitNumber := 2
	viaRunes(input, bitNumber)
	viaXOR(input, bitNumber)
}
