/*Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2....) с использованием конкурентных вычислений.*/

package task3

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func task3WaitGroupMutex(testArray []int) {
	fmt.Println("WaitGroup Mutex:")
	start := time.Now()
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	var result int
	for _, value := range testArray {
		wg.Add(1)
		go func(value int, result *int, mutex *sync.Mutex) {
			defer wg.Done()
			mutex.Lock()
			*result += value * value
			fmt.Printf("value = %v, result = %v\n", value, *result)
			mutex.Unlock()
		}(value, &result, &mutex)
	}
	wg.Wait()
	fmt.Printf("final result: %v \n", result)
	fmt.Printf("time spent: %vμs\n\n", time.Now().Sub(start).Microseconds())
}

func task3BufferedChannel(testArray []int) {
	fmt.Println("Buffered Channel:")
	start := time.Now()
	ch := make(chan int, 5)
	for _, value := range testArray {
		go func(value int) {
			ch <- value * value
		}(value)
	}
	var result int
	for i := 0; i < len(testArray); i++ {
		result += <-ch
		fmt.Printf("result: %v\n", result)
	}
	fmt.Printf("final result: %v\n", result)
	fmt.Printf("time spent: %vμs\n\n", time.Now().Sub(start).Microseconds())
}

func task3Atomic(testArray []int) {
	fmt.Println("Atomic:")
	start := time.Now()
	var (
		result uint32
		wg     sync.WaitGroup
	)

	for _, value := range testArray {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			atomic.AddUint32(&result, uint32(value*value))
			fmt.Printf("value = %v, result = %v\n", value, atomic.LoadUint32(&result))
		}(value)
	}
	wg.Wait()
	fmt.Printf("final result: %v\n", atomic.LoadUint32(&result))
	fmt.Printf("time spent: %vμs\n\n", time.Now().Sub(start).Microseconds())
}

func Task3() {
	testArray := []int{2, 4, 6, 8, 10}
	task3WaitGroupMutex(testArray)
	task3BufferedChannel(testArray)
	task3Atomic(testArray)
}
