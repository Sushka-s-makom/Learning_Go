package concurrency

import (
	"fmt"
	"time"
)

func Worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func MainSync() {
	done := make(chan bool, 1)
	go Worker(done)

	<-done
}
