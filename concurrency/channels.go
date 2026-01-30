package concurrency

import "fmt"

func Chan() {
	messages := make(chan string)
	go func() { messsages <- "ping" }()

	msg := <-messages

	fmt.Println(msg)
}
