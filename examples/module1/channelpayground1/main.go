package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int,100)
	go func() {
		fmt.Println("hello from go routine")
		for i := 0; i < 100; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(100)
			fmt.Println("Putting, ", n)
			ch <- n
		}
		close(ch)
		//ch <- 0
		//ch <- 1
	}()

	//i := <-ch
	//j := <-ch
	//fmt.Println("channel output: ", i)
	//fmt.Println("channel output: ", j)

	fmt.Println("hello from main")
	for v := range ch {
		fmt.Println("Receiving: ", v)
	}
}
