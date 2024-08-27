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
	// Метод выводит строку с информацией о человеке
	fmt.Printf("Hi, I'm %s, I'm %d years old!\n", h.Name, h.Age)
}

type Action struct {
	Human // Встраивание структуры Human
}

func (a *Action) walk() {
	// Метод выводит строку с информацией о действии
	fmt.Printf("%s, is walking.\n", a.Name)
}

func main() {
	// Создание экземпляра структуры Action
	a := Action{
		Human{
			Name: "Bob",
			Age:  25,
		},
	}

	// Так как Action встраивает Human, метод saySomething() доступен и для Action
	a.saySomething()
	a.walk()
}
