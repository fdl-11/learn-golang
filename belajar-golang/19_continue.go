package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {	// i%2 == 1 ganjil
			continue
		}

		fmt.Println("Perulangan ke", i)
	}
}