package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Fadli",
		"address": "Sragen",
	}

	person["title"] = "Ordinary Human"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// =================================================================

	var book map[string]string = make(map[string]string)
	book["title"] = "The Go Programming Language"
	book["author"] = "Fadli"
	book["ups"] = "Salah"
	
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}