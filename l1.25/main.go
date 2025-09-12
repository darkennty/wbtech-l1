package main

import (
	"fmt"
	"time"
)

func sleep(duration int) {
	timer := time.After(time.Duration(duration) * time.Millisecond)

	for {
		select {
		case <-timer:
			return
		}
	}
}

func main() {
	fmt.Println("start")

	sleep(3000) // время указано в миллисекундах

	fmt.Println("end")
}
