package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
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

	resultChan := make(chan int)
	timer := time.After(time.Duration(N) * time.Second)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(resultChan)

		i := 0
		for {
			select {
			case <-timer:
				fmt.Println("timed out")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- i
			}
			i++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case value, ok := <-resultChan:
				if !ok {
					return
				}
				fmt.Println(value)
			}
		}
	}()

	wg.Wait()
	fmt.Println("done")
}
