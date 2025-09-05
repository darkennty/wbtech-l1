package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}

	for _, v := range arr {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			i *= i
			fmt.Println(i)
		}(v)
	}

	wg.Wait()
	fmt.Println("done")
}
