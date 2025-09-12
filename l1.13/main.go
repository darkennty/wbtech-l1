package main

import "fmt"

func main() {
	a := 3
	b := 25
	fmt.Printf("Old a: %d\tOld b: %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("New a: %d\tNew b: %d\n", a, b)
}
