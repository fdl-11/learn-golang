package main

import "fmt"

func getGoodBye(name string) string {	// string kedua adalah tipe data return valuenya
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Fadli")
	fmt.Println(result)
	fmt.Println(getGoodBye("Darusalam"))
}