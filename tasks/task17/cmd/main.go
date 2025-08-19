package main

import "fmt"

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.
// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

func binSearch(arr []int, number int) int {

	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == number {
			return mid
		}
		if arr[mid] < number {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	//если после всех итераций, число не найдено, следовательно оно отсутствует в массиве
	return -1
}

func main() {
	sortedArr := []int{1, 3, 5, 23, 53, 123, 642, 2222}

	fmt.Println("Позиция числа 642 в массиве:", binSearch(sortedArr, 642))
	fmt.Println("Позиция числа 3 в массиве:", binSearch(sortedArr, 3))
	fmt.Println("Позиция числа 53 в массиве:", binSearch(sortedArr, 53))

	fmt.Println("Проверим наличие несуществующего эл-та 11 в массиве, результат поиска: ", binSearch(sortedArr, 11)) // поиск не сущ-ющего результат -1

}
