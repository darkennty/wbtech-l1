package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run l1.8/main.go <number> <i-th bit> <bit value>")
		return
	}

	var N int64
	N, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("Wrong type of an argument <number> (should be int).")
		return
	}

	i, err := strconv.Atoi(os.Args[2])
	if err != nil || i == 0 {
		fmt.Println("Wrong type of an argument <i-th bit> (should be int >= 1).")
		return
	}
	i--

	val, err := strconv.Atoi(os.Args[3])
	if err != nil || (val != 0 && val != 1) {
		fmt.Println("Wrong type of an argument <bit value> (should be 0 or 1).")
		return
	}

	var change int64
	change = 1 << i

	fmt.Printf("Old N: %d.\tBinary N: %b\n", N, N)

	if val == 0 {
		change = ^change
		N = N & change
	} else {
		N = N ^ change
	}

	fmt.Printf("New N: %d.\tBinary N: %b\n", N, N)
}

/* Examples:

$ go run l1.8/main.go 5 1 0
Old N: 5.       Binary N: 101
New N: 4.       Binary N: 100

$ go run l1.8/main.go 4 1 1
Old N: 4.       Binary N: 100
New N: 5.       Binary N: 101

$ go run l1.8/main.go 4 4 1
Old N: 4.       Binary N: 100
New N: 12.      Binary N: 1100

$ go run l1.8/main.go 5 3 0
Old N: 5.       Binary N: 101
New N: 1.       Binary N: 1

*/
