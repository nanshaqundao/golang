package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}

	fmt.Println(sum)
	name := flag.String("name", "world", "specify the name you want to say hi!")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString2 := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString2)

}
