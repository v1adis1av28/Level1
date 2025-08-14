package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
// Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
type Human struct {
	Height int
	Weight int
}

type Action struct {
	ActionName string
	Human
}

func (h *Human) Live() {
	fmt.Println("Human living!")
}

func (a *Action) GetAction() {
	fmt.Println("Human params :", a.Height, " ", a.Weight)
	a.Live()
}

func main() {
	testHuman := &Human{Height: 200, Weight: 200}
	testAction := &Action{Human: *testHuman, ActionName: "live!"}

	fmt.Println("Human do action")
	testHuman.Live()

	fmt.Println("action test on using human method")
	testAction.GetAction()
}
