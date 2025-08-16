package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
//  а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.

func Write(ch chan string) {
	for {
		ch <- "Work in progress"
	}

}

func main() {
	ch := make(chan string)
	defer close(ch)
	go Write(ch)

	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		case <-timer.C:
			fmt.Println("Program stopped with timeout")
			return
		}
	}
}
