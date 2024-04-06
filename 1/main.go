package main

import "fmt"

// определили структуру human
type Human struct {
	Name  string
	Phone string
}

//определили родительскую структуру action
type Action struct {
	Human
}

// метод для структуры human
func (h *Human) Call() {
	fmt.Println("I'm calling you from the number", h.Phone)
}

// метод для структуры action
func (a *Action) Greeting() {
	fmt.Println("Hi! I'm calling you right now!")
}

func main() {
	// создаем объект типа action
	action := Action{
		Human: Human{
			Name:  "Alexander",
			Phone: "79183746384",
		},
	}
	// через этот объект action доступен и метод call и метод greeting
	action.Call()
	action.Greeting()
}
