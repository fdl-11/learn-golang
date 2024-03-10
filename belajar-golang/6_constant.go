package main

import "fmt"

func main() {
	// Constant tidak error walaupun tidak digunakan, karena bisa saja akan digunakan di tempat lain
	const firstName string = "Fadli"
	const lastName = "Darusalam"
	const value = 3000

	const (
		age     = 21
		address = "Sragen"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	fmt.Println(age)
	fmt.Println(address)

}