package main

import "fmt"

func getFullName() (firstName string, lastName string, address string) {
	firstName = "Fadli"
	lastName = "Darusalam"
	address = "Sragen"

	return
}

func main() {
	a, b, c := getFullName()

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
}