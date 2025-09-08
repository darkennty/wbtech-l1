package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	toProcess := make(chan int)
	processed := make(chan int)

	randArr := make([]int, 0)
	for i := 0; i < 25; i++ {
		randArr = append(randArr, rand.Intn(100))
	}

	wg.Add(1)
	go func(arr []int, toProcess chan<- int) {
		defer wg.Done()

		for _, x := range arr {
			toProcess <- x
			time.Sleep(100 * time.Millisecond)
		}

		close(toProcess)
	}(randArr, toProcess)

	wg.Add(1)
	go func(ctx context.Context, toProcess <-chan int, processed chan<- int) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				close(processed)
				return
			case x, ok := <-toProcess:
				if !ok {
					close(processed)
					return
				}
				processed <- x * 2
			}
		}
	}(ctx, toProcess, processed)

	wg.Add(1)
	go func(processed <-chan int) {
		defer wg.Done()

		for {
			select {
			case value, ok := <-processed:
				if !ok {
					return
				}
				fmt.Println(value)
			}
		}
	}(processed)

	wg.Wait()
	fmt.Println("done")
}
