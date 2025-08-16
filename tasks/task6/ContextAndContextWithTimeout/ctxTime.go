package contextandcontextwithtimeout

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped context timeout expired")
			return
		default:
			fmt.Println("Worker work")
			time.Sleep(time.Second)
		}
	}
}

func exec() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting")
}
