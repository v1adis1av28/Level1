package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {
	//str := "главрыба"
	str := "☠️👽🤖"
	runSlice := []rune(str)
	reverseRune := make([]rune, len(runSlice))
	for i := len(runSlice) - 1; i >= 0; i-- {
		reverseRune = append(reverseRune, runSlice[i])
	}

	fmt.Println(string(reverseRune))
}
