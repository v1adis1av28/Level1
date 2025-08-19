package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.
// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

type Ticker struct {
	Count int
	Mu    sync.Mutex
}

func (t *Ticker) Increment() {
	t.Mu.Lock()
	t.Count++
	t.Mu.Unlock()
}

func (t *Ticker) GetCountVal() {
	t.Mu.Lock()
	fmt.Println(t.Count)
	t.Mu.Unlock()
}

func main() {
	ticker := Ticker{}
	var wg sync.WaitGroup

	for i := 0; i < 130; i++ {
		wg.Add(1)
		go func(gorutineNumber int) {
			ticker.Increment()
			wg.Done()
		}(i)
	}
	wg.Wait()
	ticker.GetCountVal()
}
