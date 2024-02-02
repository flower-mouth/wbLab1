//Реализовать все возможные способы остановки выполнения горутины.

package task6

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func goroutineOne(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Goroutine stopped via Channel Data Input")
			return
		default:
			fmt.Println("Goroutine running.")
			time.Sleep(time.Second)
		}
	}
}

func goroutineCancelViaStopChannel() {
	fmt.Println("\nCancel Via Stop Channel")
	stopChannel := make(chan bool)
	go goroutineOne(stopChannel)
	time.Sleep(3 * time.Second)
	stopChannel <- true
	time.Sleep(time.Second)
}

func goroutineTwo(stopChannel chan int) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Goroutine stopped due to Channel Closure.")
			return
		default:
			fmt.Println("Goroutine running.")
			time.Sleep(time.Second)
		}
	}
}

func goroutineCancelViaChannelClosure() {
	fmt.Println("\nCancel Via Channel Closure")
	stopChannel := make(chan int)
	go goroutineTwo(stopChannel)
	time.Sleep(3 * time.Second)
	close(stopChannel)
	time.Sleep(time.Second)
}

func goroutineThree(sharedVariable *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if *sharedVariable {
			fmt.Println("Goroutine stopped via Shared Variable.")
			return
		}

		fmt.Println("Goroutine running.")
		time.Sleep(time.Second)
	}
}

func goroutineCancelViaSharedVariable() {
	fmt.Println("\nCancel Via Shared Variable")
	var sharedVariable bool
	var wg sync.WaitGroup

	wg.Add(1)
	go goroutineThree(&sharedVariable, &wg)
	time.Sleep(3 * time.Second)
	sharedVariable = true
	wg.Wait()
}

func goroutineFour(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped via Context Cancel.")
			return
		default:
			fmt.Println("Goroutine running.")
			time.Sleep(time.Second)
		}
	}
}

func goroutineCancelViaContext() {
	fmt.Println("\nCancel Via Context")
	ctx, cancel := context.WithCancel(context.Background())
	go goroutineFour(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func goroutineFive() {
	defer fmt.Println("Goroutine stopped via Goexit.")
	for {
		fmt.Println("Goroutine running.")
		time.Sleep(time.Second)
		runtime.Goexit()
	}
}

func goroutineCancelViaGoExit() {
	fmt.Println("\nCancel Via Goexit")
	go goroutineFive()
	time.Sleep(3 * time.Second)
}

func Task6() {
	goroutineCancelViaStopChannel()
	goroutineCancelViaChannelClosure()
	goroutineCancelViaSharedVariable()
	goroutineCancelViaContext()
	goroutineCancelViaGoExit()
}
