package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Say() {
	fmt.Printf("Hello, my name is %s. I'm %d years old.\n", h.Name, h.Age)
}

type Action struct {
	Human
	Profession string
	Age        int
}

func (a *Action) Say() {
	fmt.Printf("My profession is %s\n", a.Profession)
}

func main() {
	human := Human{
		Name: "Nikita",
		Age:  21,
	}

	action := Action{
		Human:      human,
		Profession: "Software engineer",
		Age:        1,
	}

	action.Human.Say()
	action.Say()

	fmt.Printf("Names: %s, %s\n", action.Name, action.Human.Name) // same
	fmt.Printf("Ages: %d, %d\n", action.Age, action.Human.Age)    // different
}
