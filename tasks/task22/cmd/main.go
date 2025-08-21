package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).
// Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.
func add(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func substrcat(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divide(a, b *big.Int) (*big.Int, error) {
	if b == big.NewInt(0) {
		return big.NewInt(-1), fmt.Errorf("You can`t delete on zero")
	}
	return new(big.Int).Div(a, b), nil
}

func main() {

	bigIntA := big.NewInt(int64(1 << 23))
	bigIntB := big.NewInt(int64(1 << 21))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println("Adding: ", add(bigIntA, bigIntB))

	fmt.Println("Subtract: ", substrcat(bigIntA, bigIntB))

	fmt.Println("Multiply: ", multiply(bigIntA, bigIntB))

	div, err := divide(bigIntA, bigIntB)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Devision: ", div)
	}

}
