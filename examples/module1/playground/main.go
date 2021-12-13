package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}

	//??
	s := []int{2, 3, 5, 7, 11, 13}
	f := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]

	s = append(s,1,1,1,1,1,1,0)

	f = f[2:5]
	printSlice(f)

	printSlice(s)

	s = s[:2]
	printSlice(s)

	s = s[1:]
	printSlice(s)


	printSlice(f)

	// Slice the slice to give it zero length.
	f = f[:3]
	printSlice(f)

	// Extend its length.
	f = f[:4]
	printSlice(f)

	// Drop its first two values.
	f = f[2:]
	printSlice(f)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

}
type IPAddr [4]byte

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	r := make([]string, len(i))

	for a,b := range i{
		r[a]=fmt.Sprint(int(b))
	}
	result := strings.Join(r, ".")
	return "[" + result + "]"
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
