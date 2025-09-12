package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, b := big.NewInt(2<<31), big.NewInt(10<<39)

	fmt.Printf("a: %d\tb: %d\n", a, b)

	fmt.Println("a + b =", a.Add(a, b)) // a = a + b
	fmt.Printf("a: %d\tb: %d\n", a, b)

	fmt.Println("a - b =", a.Sub(a, b)) // a = a - b
	fmt.Printf("a: %d\tb: %d\n", a, b)

	fmt.Println("b / a =", b.Div(b, a)) // b = b / a
	fmt.Printf("a: %d\tb: %d\n", a, b)

	fmt.Println("b * a =", b.Mul(a, b)) // b = b * a
	fmt.Printf("a: %d\tb: %d\n", a, b)
}
