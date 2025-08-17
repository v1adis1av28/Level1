package main

import (
	"fmt"
	"sync"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

var mu sync.Mutex

func ConcurrentWrite(mp map[string]int, val int, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	mp["CW"] = val
	fmt.Println("Map key: CW  value:", mp["CW"])
	mu.Unlock()

}

func main() {

	var wg sync.WaitGroup
	mp := make(map[string]int)
	mp["CW"] = -1
	for i := 0; i < 130; i++ {
		wg.Add(1)
		go ConcurrentWrite(mp, i, &wg)

	}

	wg.Wait()
	fmt.Println("stop programm")
}
