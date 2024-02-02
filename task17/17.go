//Реализовать бинарный поиск встроенными методами языка.

package task17

import "fmt"

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (left + right) / 2
		/*fmt.Printf("Left = %v | %v\n", left, arr[left])
		fmt.Printf("Right = %v | %v\n", right, arr[right])
		fmt.Printf("Middle = %v | %v\n\n", middle, arr[middle])*/
		if target == arr[middle] {
			return middle
		} else if right-left == 1 && target == arr[right] {
			return right
		} else if target < arr[middle] {
			right = middle - 1
		} else if target > arr[middle] {
			left = middle + 1
		}
	}
	return -1
}

func Task17() {
	arr := []int{1, 6, 8, 11, 19, 25, 28, 45, 89, 104, 229, 301, 504, 669, 999, 1024, 2048, 4096, 8888, 9999}
	targetOne := 301
	targetTwo := 666
	fmt.Printf("Array:%v\n", arr)
	fmt.Printf("Target 1:%v\n", targetOne)

	result := binarySearch(arr, targetOne)
	if result != -1 {
		fmt.Printf("Element %v exists in array. It`s index = %v\n", targetOne, result)
	} else {
		fmt.Printf("Element %v doesn`t exist in array.\n", targetOne)

	}

	fmt.Printf("\n\nArray:%v\n", arr)
	fmt.Printf("Target 2:%v\n", targetTwo)

	result = binarySearch(arr, targetTwo)
	if result != -1 {
		fmt.Printf("Element %v exists in array. It`s index = %v", targetTwo, result)
	} else {
		fmt.Printf("Element %v doesn`t exist in array.", targetTwo)

	}
}
