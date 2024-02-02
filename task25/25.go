//Реализовать собственную функцию sleep.

package task25

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func Task25() {
	fmt.Println("Start of program. Sleep after 4 seconds")
	start := time.Now()
	Sleep(4 * time.Second)
	fmt.Printf("Time has passed: %v", time.Now().Sub(start).Seconds())
}
