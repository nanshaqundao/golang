package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}
type ErrNegativeSqrt float64

func (e *MyError) Error() string {
	return fmt.Sprintf("atsss %v, %s, this Error() is called",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err2 := run(); err2 != nil {
		fmt.Println(err2)
	}
	abc := ErrNegativeSqrt(-10.01)
	fmt.Println(abc)

	ddd := &MyError{time.Now(),"how old are you"}
    fmt.Println(ddd)

	i, err := strconv.Atoi("£")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

}
