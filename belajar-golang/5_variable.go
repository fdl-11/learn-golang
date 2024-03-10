package main

import "fmt"

func main() {
    var name string

	name = "Fadli Darusalam"
    fmt.Println(name)

	name = "Fadli"
	fmt.Println(name)

	// =================================================================

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// =================================================================

	country := "Indonesia"	//":="" adalah pengganti var (hanya untuk inisiasi awal)
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)

	// =================================================================

	var (
		firstName = "Fadli"
		lastName  = "Darusalam"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}