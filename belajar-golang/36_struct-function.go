package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	fadli := Customer{
		Name: "Fadli",
        Age: 30,
	}

	fadli.sayHello()
}