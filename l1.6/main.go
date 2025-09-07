package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Выход по условию
	wg.Add(1)
	timer := time.After(3 * time.Second)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-timer:
				fmt.Println("Goroutine #1 stopped.")
				return
			default:
				fmt.Println("Goroutine #1 working...")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()

	// Выход с помощью контекста
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine #2 stopped.")
				return
			default:
				fmt.Println("Goroutine #2 working...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()

	// Выход через канал уведомления
	notify := make(chan int, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-notify:
				fmt.Println("Goroutine #3 stopped.")
				return
			default:
				fmt.Println("Goroutine #3 working...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	notify <- 1
	wg.Wait()
	close(notify)

	// Выход через закрытие канала
	closing := make(chan int, 3)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case _, ok := <-closing:
				if !ok {
					fmt.Println("Goroutine #4 stopped.")
					return
				}

				fmt.Println("Goroutine #4 working...")
				time.Sleep(time.Second)
			}
		}
	}()

	closing <- 1
	closing <- 2
	closing <- 3
	close(closing)
	wg.Wait()

	// Выход через runtime.Goexit()
	timer = time.After(3 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-timer:
				stop()
			default:
				fmt.Println("Goroutine #5 working...")
				time.Sleep(time.Second)
			}
		}
	}()

	wg.Wait()
	fmt.Println("done")
}

func stop() {
	fmt.Println("Goroutine #5 stopped.")
	runtime.Goexit()
}
