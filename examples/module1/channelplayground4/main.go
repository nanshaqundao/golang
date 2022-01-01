package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(ctx)
	//select {
	//case <-ctx.Done():
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("main process exit!")
	//}

	time.Sleep(time.Second * 10)
}
