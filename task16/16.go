/*Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.*/

package task16

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	} else {
		pivot := arr[0]
		var (
			left  []int
			right []int
			equal []int
		)
		for _, value := range arr {
			if value < pivot {
				left = append(left, value)
			}
			if value == pivot {
				equal = append(equal, value)
			}
			if value > pivot {
				right = append(right, value)
			}
		}
		/*		fmt.Printf("Array:%v\n", arr)
				fmt.Printf("Pivot:%v\n", pivot)
				fmt.Printf("Left:%v\n", left)
				fmt.Printf("Equal:%v\n", equal)
				fmt.Printf("Right:%v\n\n", right)*/

		return append(append(quickSort(left), equal...), quickSort(right)...)
	}
}

func Task16() {
	var arr = []int{5, 6, 7, 10, 99, 52, 62, 999, 512, 256, 0, 1, 55, 4, 3, 1024, 88, 66}
	fmt.Printf("Given array:%v\n", arr)
	fmt.Printf("Sorted array%v\n", quickSort(arr))
}
