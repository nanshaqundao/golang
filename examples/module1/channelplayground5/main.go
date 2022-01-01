package main

import (
	"fmt"
	"time"
)

func main() {
	// producer put to channel every second but consumer read every 2 seconds
	// after few seconds, the channel is blocking until consumer read


	ch := make(chan int, 10)
	ticker := time.NewTicker(1 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)
	// producer
	go func() {
		flagInt := 0
		fmt.Println("hello from producer")
		for _ = range ticker.C {
			flagInt = flagInt + 1
			fmt.Println("put integer to channel: ", flagInt)
			ch <- flagInt
		}
		//ch <- 0
		//ch <- 1
	}()


	// consumer
	fmt.Println("hello from consumer - main")
	for _ = range ticker2.C {

		select {
		case v:= <- ch:

			fmt.Println("Receiving: ", v)

		default:
			fmt.Println("dumb step")
		}
	}

}
