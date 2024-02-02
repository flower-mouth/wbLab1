/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).*/

package task1

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
	Height float64
}

func (h *Human) changeAge(newValue int) {
	h.Age = newValue
}

func (h *Human) changeName(newValue string) {
	h.Name = newValue
}

func (h *Human) getData() {
	fmt.Printf("Age: %v, Name: %v\n", h.Age, h.Name)
}

func (a *Action) changeHeight(newValue float64) {
	a.Height = newValue
}

func (a *Action) getData() {
	fmt.Printf("Age: %v, Name: %v, Height: %v\n", a.Age, a.Name, a.Height)
}

func Task1() {
	JohnDoe := Human{
		Age:  22,
		Name: "John Doe",
	}
	JaneDoe := Action{
		Human:  Human{20, "Jane Doe"},
		Height: 175,
	}
	JohnDoe.getData()
	JohnDoe.changeAge(35)
	JohnDoe.changeName("John Doe Jr")
	JohnDoe.getData()
	JaneDoe.getData()
	JaneDoe.changeAge(18)
	JaneDoe.changeName("Jane Doe II")
	JaneDoe.changeHeight(179)
	JaneDoe.getData()
}
