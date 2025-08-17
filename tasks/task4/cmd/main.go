package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров
// при получении сигнала прерывания.
// Подсказка: можно использовать контекст (context.Context) или
// канал для оповещения о завершении.

func worker(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stoping workers!")
			return
		default:
			fmt.Println("we working")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go worker(&wg, ctx)

	<-sigCh

	stop()

	wg.Wait()
	fmt.Println("Stoping gorutine gracefully")
}
