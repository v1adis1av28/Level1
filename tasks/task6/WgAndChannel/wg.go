package wgandchannel

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, exitChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-exitChan:
			fmt.Printf("Worker %d shutting down...\n", id)
			return
		default:
			fmt.Printf("Worker %d running...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	exitChan := make(chan struct{})

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, exitChan)
	}

	time.Sleep(3 * time.Second)
	close(exitChan)
	wg.Wait()
	fmt.Println("All workers exitin")
}
