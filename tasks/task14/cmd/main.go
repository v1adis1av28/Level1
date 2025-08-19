package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

func checkType(t interface{}) {
	fmt.Println("Определим тип переменной:")

	switch value := t.(type) {
	case int:
		fmt.Println("Переменная равна типу int и ее значение: ", value)
	case string:
		fmt.Println("Переменная равна типу string и ее значение: ", value)
	case bool:
		fmt.Println("Переменная равна типу bool и ее значение: ", value)
	default:
		if reflect.TypeOf(t).Kind() == reflect.Chan { //Используем рефлект чтобы была вариативность и мы могли определять что пременная является каналом любого типа
			fmt.Println("Переменная равна типу chan и ее значение: ", value)
		} else {
			fmt.Println("Не удалось определить тип вашей перемнной")
		}
	}
}

func main() {
	testInt := 13
	testString := "test"
	testBool := true
	testChan := make(chan string)
	unknownType := 13.0 // Сделаем дополнительно тип который не указан и в обработчик добавим поведение при обнаружении неизвестного типа

	checkType(testInt)
	checkType(testString)
	checkType(testBool)
	checkType(testChan)
	checkType(unknownType)
}
