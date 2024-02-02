/*Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.*/

package task4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(ch chan int, number int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nWorker %v stopped.", number)
			return
		case data := <-ch:
			fmt.Printf("Worker %v got: %v\n", number, data)
		}
	}
}

func Task4() {
	var (
		quantity     int
		channelInput int
	)

	ch := make(chan int)

	fmt.Println("Enter number of workers: ")
	fmt.Scan(&quantity)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	for i := 1; i <= quantity; i++ {
		go worker(ch, i, ctx)
	}

	for {
		select {
		case <-ctx.Done():
			close(ch)
			os.Exit(100)
		default:
			ch <- channelInput
			channelInput++
			time.Sleep(300 * time.Millisecond)
		}
	}
}
