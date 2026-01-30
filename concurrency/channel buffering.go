package concurrency

import "fmt"

func ChanBuff() {
	messages := make(chan string, 3)
	messages <- "hello"
	messages <- "world"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
