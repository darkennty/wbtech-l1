package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // Создаётся строка длиной 1024 символа
	//justString = v[:100] Переменная ссылалась бы на большую по размеру переменную v, из-за чего v оставалась бы в памяти

	temp := make([]byte, 100)
	copy(temp, v[:100]) // Копируем 100 элементов из большой строки (избавляемся от ссылки на v)

	justString = string(temp) // Теперь нет ссылки на v - память, занимаемая v, будет очищена сборщиком мусора
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	return strings.Repeat("a", size) // Создаём строку размером size
}
