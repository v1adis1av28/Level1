package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
// Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

func Square(val int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- val * val
}

func main() {
	testArr := []int{2, 4, 6, 8, 10}
	arCh := make(chan int, len(testArr))
	var wg sync.WaitGroup
	wg.Add(len(testArr))
	for _, val := range testArr {
		go Square(val, arCh, &wg)
	}

	wg.Wait()
	close(arCh)
	ansArr := make([]int, 0, len(testArr))
	for val := range arCh {
		ansArr = append(ansArr, val)
	}
	fmt.Println(ansArr)
}
