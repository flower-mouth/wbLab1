//Реализовать конкурентную запись данных в map.

package task7

import (
	"fmt"
	"sync"
	"time"
)

func withSyncMutex() {
	fmt.Println("\n\nConcurrent writing in map with sync.Mutex")
	start := time.Now()
	m := make(map[int]string)
	var (
		mutex sync.Mutex
		wg    sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			mutex.Lock()
			m[key] = fmt.Sprintf("value of key %v", key)
			mutex.Unlock()
			fmt.Printf("key: %v\t  value: %v\n", key, m[key])
		}(i)
	}
	wg.Wait()
	fmt.Printf("time spent: %vs\n", time.Now().Sub(start).Seconds())
}

func withSyncMap() {
	fmt.Println("\n\nConcurrent writing in map with sync.Map")
	start := time.Now()
	var (
		m  sync.Map
		wg sync.WaitGroup
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			m.Store(key, fmt.Sprintf("value of key %v", key))
			value, _ := m.Load(key)
			fmt.Printf("key: %v\t  value: %v\n", key, value)
		}(i)
	}
	wg.Wait()
	fmt.Printf("time spent: %vs\n", time.Now().Sub(start).Seconds())
}

func Task7() {
	withSyncMutex()
	withSyncMap()
}
