package main

import "fmt"

// Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.

func quickSort(arr []int, low, high int) []int {
	var pivot int
	if low < high {
		arr, pivot = partition(arr, low, high)
		arr = quickSort(arr, low, pivot-1)
		arr = quickSort(arr, pivot+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func main() {
	testArr := []int{3, 1, 10}
	fmt.Println(quickSort(testArr, 0, len(testArr)-1))
}
