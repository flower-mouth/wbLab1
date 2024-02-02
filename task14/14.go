/*Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.*/

package task14

import "fmt"

func Task14() {
	var s = []interface{}{0, "str", false, 99 - 2i, make(chan int), struct{}{}, make(chan int64)}

	for _, val := range s {
		fmt.Printf("%T: %v\n", val, val)
	}

	fmt.Println()

	for _, val := range s {
		switch v := val.(type) {
		case int:
			fmt.Printf("Integer: %v\n", v)
		case float64:
			fmt.Printf("Float64: %v\n", v)
		case string:
			fmt.Printf("String: %v\n", v)
		case bool:
			fmt.Printf("Boolean: %v\n", v)
		case chan int:
			fmt.Printf("Chan int: %v\n", v)
		default:
			fmt.Printf("Unknown type\n")
		}
	}
}
