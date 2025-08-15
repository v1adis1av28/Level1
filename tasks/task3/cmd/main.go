package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

func WriteToChannel(ch chan string) {
	ch <- "Write to channel"
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("You should add number of workers as args in CLI")
	}
	N, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal("Error on parsing int from cli")
	}
	ch := make(chan string)
	for i := 0; i < N; i++ {
		go ReadFromChannel(i, ch)
	}
	for {
		go WriteToChannel(ch)
	}

}

func ReadFromChannel(i int, ch chan string) {
	for {
		fmt.Println("Read from goroutine number :", i, " ", <-ch)
	}

}
