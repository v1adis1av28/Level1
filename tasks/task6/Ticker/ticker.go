package ticker

import (
	"fmt"
	"sync"
	"time"
)

func Work(ch chan struct{}, wg *sync.WaitGroup, id int) {
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-ch:
			wg.Done()
			fmt.Println("exit")
			return
		case <-ticker.C:
			fmt.Println("Work")
		}
	}

}

func exec() {
	var wg sync.WaitGroup
	c := make(chan struct{}, 1)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Work(c, &wg, i)
	}
	<-time.After(time.Second * 2)
	close(c)
	fmt.Println("waiting")
	wg.Wait()

	fmt.Println("Done")

}
