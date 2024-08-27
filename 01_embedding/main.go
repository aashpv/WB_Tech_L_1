package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
}

func (h *Human) saySomething() {
	fmt.Printf("Hi, I'm %s, I'm %d years old!\n", h.Name, h.Age)
}

type Action struct {
	Human
}

func (a *Action) walk() {
	fmt.Printf("%s, is comming.\n", a.Name)
}

func main() {
	a := Action{
		Human{
			Name: "Bob",
			Age:  25,
		},
	}
	a.saySomething()
	a.walk()
}
