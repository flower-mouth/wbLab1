/*Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.*/

package task18

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type counter struct {
	value int
	mx    sync.Mutex
}

func (T *counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	T.mx.Lock()
	T.value++
	T.mx.Unlock()
}

type counterAtomic struct {
	value int32
}

func (T *counterAtomic) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&T.value, 1)
}

func viaMutex() {
	fmt.Printf("\nViaMutex:\n")
	start := time.Now()
	var (
		cnt counter
		wg  sync.WaitGroup
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go cnt.Increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Final value = %v\n", cnt.value)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}

func viaAtomic() {
	fmt.Printf("\nViaAtomic:\n")
	start := time.Now()
	var (
		cntAtomic counterAtomic
		wg        sync.WaitGroup
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go cntAtomic.Increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Final value = %v\n", cntAtomic.value)
	fmt.Printf("time spent: %vms\n\n", time.Now().Sub(start).Milliseconds())
}

func Task18() {
	viaMutex()
	viaAtomic()
}
