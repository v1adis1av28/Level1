package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.
// Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
// Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

func Sleep(duration int) {
	<-time.After(time.Duration(duration) * time.Second)
}

func main() {
	duration := 5
	start := time.Now()

	Sleep(duration)

	end := time.Since(start)
	fmt.Println(end)
}
