/*Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.*/

package task5

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func writer(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for i := 0; i < 500; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			time.Sleep(500 * time.Millisecond)
		}
	}
	close(ch)
}

func reader(ch chan int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time exceeded. Exiting.")
			return
		case data, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Data from channel: %v\n", data)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func Task5() {
	start := time.Now()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go writer(ch, &wg, ctx)
	go reader(ch, &wg, ctx)

	wg.Wait()
	fmt.Printf("time spent: %vs\n\n", time.Now().Sub(start).Seconds())
}
