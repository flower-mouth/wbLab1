/*Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.*/

package task2

import (
	"fmt"
	"sync"
	"time"
)

func task2WaitGroup(testArray []int) {
	var wg sync.WaitGroup
	for _, value := range testArray {
		wg.Add(1)
		go func(value int) {
			fmt.Printf("%v\t", value*value)
			wg.Done()
		}(value)
	}
	wg.Wait()
}

func task2BufferedChannel(testArray []int) {
	ch := make(chan int, 5)
	defer close(ch)

	for _, value := range testArray {
		go func(value int) {
			ch <- value * value
		}(value)
	}
	for i := 0; i < len(testArray); i++ {
		fmt.Printf("%v\t", <-ch)
	}
}

func task2TimeSleep(testArray []int) {
	for _, value := range testArray {
		go func(value int) {
			fmt.Printf("%v\t", value*value)
		}(value)
	}
	time.Sleep(50 * time.Millisecond)
}

func Task2() {
	testArray := []int{2, 4, 6, 8, 10}
	fmt.Println("WaitGroup:")
	task2WaitGroup(testArray)
	fmt.Println("\nBuffered Channel:")
	task2BufferedChannel(testArray)
	fmt.Println("\nTime Sleep 100ms:")
	task2TimeSleep(testArray)
}
