package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments to start program.")
		return
	}

	str := []rune(os.Args[1])
	newStr := make([]rune, len(str))

	k := 0
	for i := len(str) - 1; i >= 0; i-- {
		newStr[k] = str[i]
		k++
	}

	fmt.Println(string(newStr))
}

/* Примеры запуска:
$ go run l1.19/main.go главрыба
абырвалг

$ go run l1.19/main.go 😂🍕🎬🎹
🎹🎬🍕😂
*/
