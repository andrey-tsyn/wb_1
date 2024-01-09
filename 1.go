package main

import "fmt"

type Human struct {
	Age     int
	Name    string
	isAlive bool
}

func (h Human) Die() {
	h.isAlive = false
}

func (h Human) IsAlive() bool {
	return h.isAlive
}

func (h Human) DoSomething() {
	fmt.Println("Do something from Human...")
}

type Action struct {
	// Встраиваем поля и методы структуры Human в структуру Action
	Human
	SomeStr string
}

func (a Action) DoSomething() {
	fmt.Println("Do something from Action...")
}

func main() {
	action := Action{
		Human: Human{
			Age:     23,
			Name:    "Andrew",
			isAlive: true,
		},
		SomeStr: "String",
	}

	fmt.Printf("Age is: %d\n", action.Age)
	fmt.Printf("Name is: %s\n", action.Name)
	fmt.Printf("Is alive: %t\n", action.IsAlive())

	// Метод DoSomething затенен в Action, поэтому метод структуры Human не выполнится
	action.DoSomething()
}
