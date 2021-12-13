package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)

	//var huma1 IF
	//huma1 = new(Human)
	//
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)

	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	interfaces = append(interfaces, p)
	fmt.Println(p.getName())


	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
}
