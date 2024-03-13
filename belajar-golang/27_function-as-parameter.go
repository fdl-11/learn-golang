package main

import "fmt"
import "strings"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)

	// sayHelloWithFilter(nameFiltered, upperFilter)
	
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Hayoloooo"
	} else {
		return name
	}
}

func upperFilter(name string) string {
	nameFiltered := strings.ToUpper(name)
	return nameFiltered
}

func main() {
	sayHelloWithFilter("Fadli", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Tes", filter)
}