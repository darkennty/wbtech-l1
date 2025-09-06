package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments to start program.")
		return
	}

	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Wrong type of an argument (should be int).")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	numbersToProcess := make(chan int, N)

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToProcess)
		}()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				numbersToProcess <- rand.Intn(100)
			}
		}
	}()

	<-quit
	cancel()
	close(numbersToProcess)
}

func worker(ctx context.Context, toProcess <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			if !ok {
				return
			}

			fmt.Println(value)
		}
	}
}
