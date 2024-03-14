package main

import "fmt"
import "strings"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	return filter(name)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Hayoloooo"
	}

	return name
}

func upperFilter(name string) string {
	return strings.ToUpper(name)
}

func main() {
	names := []string{"Fadli", "Anjing"}

	for _, name := range names {
		fmt.Println(name)
		fmt.Println(sayHelloWithFilter(name, spamFilter))
		fmt.Println(sayHelloWithFilter(sayHelloWithFilter(name, spamFilter), upperFilter))
		fmt.Println()
	}

}