package main

import (
	"fmt"
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
//  После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
//  Убедитесь, что чтение из второго канала корректно завершается.

func ReadNumber(in chan int, out chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func DoubleNumber(out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range out {
		fmt.Println(i)
	}
}

func main() {

	var wg sync.WaitGroup

	in := make(chan int)
	out := make(chan int)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 1, 13, 14, 15, 16, 16, 71, 15, 43, 234, 5, 123, 1, 23, 55, 35, 654, 2, 213, 234, 123}
	go ReadNumber(in, out)
	wg.Add(1)

	go DoubleNumber(out, &wg)
	for _, val := range arr {
		in <- val
	}
	close(in)
	wg.Wait()

	fmt.Println("Programm stoped")
}
