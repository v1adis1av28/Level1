package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.
// Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.
// Подумайте, какой структурой данных удобно воспользоваться для проверки условия.

func checkUniqueString(str string) bool {
	unique := make(map[rune]bool)
	lowStr := strings.ToLower(str)
	for _, val := range lowStr {
		if unique[val] == true {
			return false
		} else {
			unique[val] = true
		}
	}
	return true
}

func main() {
	strTrue := "abcd"
	strFalse := "abCdefAaf"

	fmt.Println(checkUniqueString(strTrue))
	fmt.Println(checkUniqueString(strFalse))
}
