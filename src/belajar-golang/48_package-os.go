/*
Package os berisi fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan di semua sistem operasi)
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args // lokasi file hasil kompilasi
    fmt.Println(args)

	name, err := os.Hostname()
	if err != nil {
        fmt.Println(err)
    } else {
		fmt.Println(name)
	}

	// export APP_USERNAME = root
	// export APP_PASSWORD = root
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username, password)
}