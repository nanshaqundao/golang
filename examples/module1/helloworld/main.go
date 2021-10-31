package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)

	// str := "hello world"
	// fmt.Printf("%d\n", str)

	// i := 3
	// fmt.Println(i != 0|| i !=1 )

	// words := []string{"foo", "bar", "baz"}
	// for _, word := range words {
	// 	go func() {
	// 		fmt.Println(word)
	// 	}()
	// }
}
