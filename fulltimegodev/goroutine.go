package main

import (
	"fmt"

	"time"
)

func channels() {
	msgCh := make(chan string, 128)

	msgCh <- "Hi"
	msgCh <- "Hello"
	msgCh <- "How're you doing"
	msgCh <- "I am good thank you"
	close(msgCh)

	for {
		msg, ok := <-msgCh
		if !ok {
			break
		}
		fmt.Println("the message is:", msg)
	}
	// this is our consumers
	for msg := range msgCh {
		fmt.Println("the msg is:", msg)
	}

	fmt.Println("Done reading from the channel")
}

func fetchData() string {
	time.Sleep(time.Second * 2)
	return "some result"
}
