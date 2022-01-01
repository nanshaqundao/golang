package main

func main() {
	ch := make(chan int, 100)
	go producerFunc(ch)
	go consumerFunc(ch)
}

func producerFunc(ch chan<- int) {
	for {
		ch <- 1
	}
}

func consumerFunc(ch <-chan int) {
	for {
		<-ch
	}
}
