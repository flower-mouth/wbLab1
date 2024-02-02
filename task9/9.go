/*Разработать конвейер чисел. Даны два канала:
в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.*/

package task9

import "fmt"

func writeToChannel(ch chan int, arr []int) {
	defer close(ch)
	for i := range arr {
		ch <- arr[i]
	}
}

func readFromChannel(chInput, chOutput chan int) {
	defer close(chOutput)
	for elem := range chInput {
		chOutput <- elem * 2
	}
}

func Task9() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	chan1 := make(chan int)
	chan2 := make(chan int)
	go writeToChannel(chan1, arr)
	go readFromChannel(chan1, chan2)

	for elem := range chan2 {
		fmt.Printf("%v, ", elem)
	}
}
