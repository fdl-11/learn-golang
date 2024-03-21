package main

import "fmt"

func main() {
	name := "Fadli"
	counter := 0

	increment := func() {	// Anonymous function
		name := "Darusalam"	// Agar tidak mereplace variabel name yg sudah ada di luar, maka perlu deklarasi ulang atau bedakan nama
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}