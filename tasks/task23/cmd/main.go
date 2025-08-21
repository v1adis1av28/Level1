package main

import "fmt"

// Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	elIndex := 3
	fmt.Println("delete el by index ", elIndex)
	fmt.Println(len(arr), cap(arr))
	arr, err := RemoveElByIndex(arr, elIndex)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("arr after deleting: ", arr)
		fmt.Println(len(arr), cap(arr))
	}
}

func RemoveElByIndex(arr []int, elIndex int) ([]int, error) {
	if elIndex < 0 || elIndex > len(arr) {
		return []int{}, fmt.Errorf("Wrong value of element index")
	}
	if len(arr) < 1 {
		return []int{}, fmt.Errorf("Nothing to remove in empty array")
	}

	return append(arr[:elIndex], arr[elIndex+1:]...), nil
}

//go run main.go
// [1 2 3 4 5 6 7 8]
// delete el by index  3
// 8 8
// arr after deleting:  [1 2 3 5 6 7 8]
// 7 8
