package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Fadli"
	names[1] = "Darusalam"
	names[2] = "Sragen"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// =================================================================

	var values = [3]int{
		90, 
		80, 
		85,
	}

	fmt.Println(values)

	// len() adalah panjang arraynya, bukan jumlah datanya
	fmt.Println(len(names))
	fmt.Println(len(values))
}