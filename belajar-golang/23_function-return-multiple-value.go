package main

import "fmt"

func getFullName() (string, string, string) {
	return "Fadli", "Darusalam", "Sragen"
}

func main() {
	firstName, _, address := getFullName()		// tanda underscore(_) untuk menghiraukan / ignore value yg diambil
	fmt.Println(firstName)
	// fmt.Println(lastName)
	fmt.Println(address)
}