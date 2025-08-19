package main

import (
	"fmt"
	"strings"
)

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

// Идея решение это проходить по последовательности и использовать идиому "ок" для проверки наличия слова в последовательности
// грубо говоря реализуем сет
func GetUniqueSequense(sq []string) []string {
	mp := make(map[string]int)
	var res []string
	for _, val := range sq {
		//Также приведем к нижнему регистру ключи, так как синтаксически "Dog" и "dog" одно слово
		_, ok := mp[strings.ToLower(val)]
		if !ok { // если такого ключа нет
			mp[val]++
			res = append(res, val)
		} else {
			continue
		}
	}
	return res
}

func main() {
	testSeq := []string{"cat", "cat", "dog", "cat", "tree", "CAt", "DOG"}
	uniqSeq := GetUniqueSequense(testSeq)

	fmt.Println(uniqSeq)
}
