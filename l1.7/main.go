package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	mapa := make(map[int]int)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()

			mu.Lock()
			mapa[rand.Intn(20)] = i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for k, v := range mapa {
		fmt.Println(k, v)
	}
	fmt.Println("done")
}
