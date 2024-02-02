/*Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.*/

package task12

import (
	"fmt"
	"time"
)

func Task12() {
	start := time.Now()
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	result := []string{}
	hashMap := make(map[string]int)
	for _, value := range arr {
		hashMap[value]++
		if hashMap[value] == 1 {
			result = append(result, value)
		}
	}
	fmt.Println(result)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}
