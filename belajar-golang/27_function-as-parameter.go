package main

import "fmt"
import "strings"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	nameFiltered := filter(name)

	return nameFiltered
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
	name1 := "Fadli"
	name2 := "Anjing"

	// Original name
	fmt.Println(name1)
	fmt.Println(name2)
	
	fmt.Println()
	
	// SpamFilter
	name1 = sayHelloWithFilter(name1, spamFilter)
	name2 = sayHelloWithFilter(name2, spamFilter)
	
	fmt.Println(name1)
	fmt.Println(name2)
	
	fmt.Println()

    // UpperFilter
	fmt.Println(sayHelloWithFilter(name1, upperFilter))
	fmt.Println(sayHelloWithFilter(name2, upperFilter))

}