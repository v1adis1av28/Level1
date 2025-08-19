package main

import (
	"fmt"
	"strings"
)

// Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?
// Приведите корректный пример реализации.
// Вопрос: что происходит с переменной justString?

var justString string

// func someFunc() {
//   v := createHugeString(1 &lt;&lt; 10) //Создаем длинную строку, но после выполнения функии данная переменная пропадет
//   justString = v[:100] // И вот здесь строка будет продолжать ссылаться на участок памяти.
//	 Получается что в памяти остается вся длинная строка, однако использоваться будут только первые 100 байт, что приводит к удержанию памяти
// } Решим данную проблему тем, что склонируем часть необходимой нам строки, вместо ссылки на саму строку и определнное кол-во байт

func main() {
	someFunc()

	fmt.Println(justString)      // сама строка
	fmt.Println(len(justString)) //100
}

func someFunc() {
	v := createHugeString(512)
	justString = strings.Clone(v[:100])
}

func createHugeString(i int) string {
	str := make([]byte, i)
	for ind := range str {
		str[ind] = 'A'
	}

	return string(str)
}
