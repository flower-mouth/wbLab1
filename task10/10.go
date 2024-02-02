/*Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

package task10

import (
	"fmt"
	"time"
)

func Task10() {
	start := time.Now()
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	hashMap := make(map[int][]float32)
	for _, value := range arr {
		buff := int(value / 10)
		hashMap[buff*10] = append(hashMap[buff*10], value)
	}
	for key, value := range hashMap {
		fmt.Printf("key = %v, value = %v\n", key, value)
	}
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}
