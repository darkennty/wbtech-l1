package main

import "fmt"

/*
	С помощью паттерна "Адаптер" можно добиться вызова у структуры с несовместимым
	интерфейсом методов, которые не описаны в данной структуре, без изменения самой структуры.
	Пример использования: одинаковое использование разных логгеров с разными функциями для лога.
*/

type Dog struct{}

func (d *Dog) Speak() {
	fmt.Println("bark-bark-bark")
}

type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("meow")
}

type AnimalAdapter interface {
	Speak()
}

type CatAdapter struct {
	*Cat
}

func (a *CatAdapter) Speak() {
	a.Cat.Meow()
}

func main() {
	dog := Dog{} // У каждой из структур свой метод подачи голоса
	cat := Cat{} // У Dog{} это метод Speak() (подходит), у Cat{} - метод Meow() (не подходит)

	animal := AnimalAdapter(&dog) // адаптер для структуры Dog
	animal.Speak()                // используем одинаковый вызов метода подачи голоса

	animal = AnimalAdapter(&CatAdapter{&cat}) // адаптер для структуры Cat
	animal.Speak()                            // используем одинаковый вызов метода подачи голоса
}
