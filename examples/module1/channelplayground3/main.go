package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for {

			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return

			default:
				fmt.Println("looping....")
			}

		}
	}()
	time.Sleep(5 * time.Second)
	//stop goroutine by channel (close will send a message to channel)
	close(done)
}
