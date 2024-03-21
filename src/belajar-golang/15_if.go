package main

import "fmt"

func main() {
	name := "Darusalam"

	if name == "Fadli" {
		fmt.Println("Halo Fadli")
    } else if name == "Darusalam" {
		fmt.Println("Halo Darusalam")
	} else {
		fmt.Println("Hi, kenalan yuk...")
	}

	// length := len(name)
	// if length > 5 {
	// 	fmt.Println("Nama terlalu panjang")
	// } else {
	// 	fmt.Println("Nama sudah sesuai")
	// }

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah sesuai")
	}

}