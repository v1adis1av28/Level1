package main

import "fmt"

// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).

func main() {
	var num int64 = 10 //1010
	fmt.Println(num)
	num = SetBit(num, 2, 1)
	fmt.Println(num)

	//вернем обратно число
	num = SetBit(num, 2, 0)
	fmt.Println(num)

	num = SetBit(num, 0, 1)
	fmt.Println(num)
}

func SetBit(num int64, index, value int) int64 {
	if value == 1 {
		return num | (1 << index)
	} else {
		return num &^ (1 << index)
	}
}
