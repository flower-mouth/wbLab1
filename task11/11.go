//Реализовать пересечение двух неупорядоченных множеств.

package task11

import (
	"fmt"
	"time"
)

func Task11() {
	start := time.Now()
	first := []int{1, 2, 5, 9, 7, 4, 3, 999, 15, 88, 69, 558}
	second := []int{558, 999, 101, 222}
	result := []int{}
	hashMap := make(map[int]int)
	for _, value := range first {
		hashMap[value]++
	}
	for _, value := range second {
		if _, ok := hashMap[value]; ok {
			result = append(result, value)
		}
	}
	fmt.Println(result)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}
