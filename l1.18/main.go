package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Используем обе реализации конкурентного счётчика, чтобы сравнить скорость выполнения
	addByMutex()
	addByAtomic()
}

func addByMutex() {
	var (
		counter int64
		wg      sync.WaitGroup
		mu      sync.Mutex
	)
	start := time.Now()
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	workTime := time.Since(start).Milliseconds()

	fmt.Println("Increment by mutex:")
	fmt.Printf("Counter: %d \t Time: %d ms\n", counter, workTime) // Работает медленнее
	fmt.Println()
}

func addByAtomic() {
	var (
		counter int64
		wg      sync.WaitGroup
	)
	start := time.Now()

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()

			atomic.AddInt64(&counter, 1)
		}()
	}

	wg.Wait()

	workTime := time.Since(start).Milliseconds()

	fmt.Println("Increment by atomic:")
	fmt.Printf("Counter: %d \t Time: %d ms\n", counter, workTime) // Работает быстрее
}

/* results:
Increment by mutex:
Counter: 1000    Time: 4 ms

Increment by atomic:
Counter: 1000    Time: 1 ms
*/
