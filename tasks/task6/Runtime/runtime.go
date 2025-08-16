package runtime

import (
	"fmt"
	"runtime"
	"time"
)

func Work() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("Start")

	runtime.Goexit()
}

func Check() {
	go Work()

	time.Sleep(2 * time.Second)
	fmt.Println("stop")
}
