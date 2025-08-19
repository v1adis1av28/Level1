package main

import "fmt"

//Поменять местами два числа без использования временной переменной./
// Подсказка: примените сложение/вычитание или XOR-обмен.

func ArithmeticSolution(a, b int) {
	fmt.Println(fmt.Sprintf("a:%d b:%d", a, b))
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Arithmetic solution")
	fmt.Println(fmt.Sprintf("a:%d b:%d", a, b))
}

func XORSolution(c, d int) {
	fmt.Println(fmt.Sprintf("c:%d d:%d", c, d))
	c = c ^ d
	d = c ^ d
	c = c ^ d
	fmt.Println("XOR solution")
	fmt.Println(fmt.Sprintf("c:%d d:%d", c, d))
}

func main() {
	a, b := 2, 3
	ArithmeticSolution(a, b)
	fmt.Println()
	c, d := 2, 3
	XORSolution(c, d)
}
