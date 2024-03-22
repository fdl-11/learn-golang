/*
Package flag berisi fungsionalitas untuk memparsing command line argument
*/

package main

import (
    "flag"
    "fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database hostname")
	user := flag.String("user", "root", "Put your database username")
	pass := flag.String("pass", "root", "Put your database password")

	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()

	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Pass: ", *pass)
	fmt.Println("Number: ", *number)

	// go run 49_package-flag.go
	// go run 49_package-flag.go -host=localhost -user=fadli
	// go run 49_package-flag.go -host=localhost -user=fadli -number=fadli (error tampil helper)

}