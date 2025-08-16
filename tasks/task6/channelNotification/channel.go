package channelnotification

import (
	"fmt"
	"time"
)

func worker(exitChan chan struct{}) {
	for {
		select {
		case <-exitChan:
			fmt.Println("Worker shutting down...")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(time.Second)
		}
	}
}

func exec() {
	exitChan := make(chan struct{})
	go worker(exitChan)

	time.Sleep(3 * time.Second)
	//Чтение из закрытого канала возавращет 0, поэтому будет выпроенн селект на 11 строке
	close(exitChan)
	time.Sleep(time.Second)
	fmt.Println("Stoping programm")
}
