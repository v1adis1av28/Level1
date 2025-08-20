package main

import "fmt"

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».
// Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func reverseRunes(runes []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
}

//сперва перевернем всю строку, потом будет итеративно проходить по ней и если встречаем пробел, то переворачиваем участок до пробела, тем самым возвращая слово в обратный порядок, но в перевернутом предложении
func reverseOrder(s string) string {
	runes := []rune(s)
	length := len(runes)

	reverseRunes(runes, 0, length-1)

	start := 0
	for i := 0; i < length; i++ {
		if runes[i] == ' ' {
			reverseRunes(runes, start, i-1)
			start = i + 1
		}
	}

	reverseRunes(runes, start, length-1)

	return string(runes)
}

func main() {
	testStr := "snow dog sun s as"
	reverseOrderStr := reverseOrder(testStr)
	fmt.Println(reverseOrderStr)
}
