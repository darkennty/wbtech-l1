package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(getType(3))
	fmt.Println(getType("Hello World!"))
	fmt.Println(getType(true))
	fmt.Println(getType(make(chan int)))
	fmt.Println(getType(make(chan string)))
	fmt.Println(getType(make(chan bool)))
}

func getType(v interface{}) string {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "can't get type of the variable"
	}
}
