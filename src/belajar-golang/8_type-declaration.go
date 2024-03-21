package main

import "fmt"

func main() {
    type NoKTP string	// tipe data alias, membuat tipe data baru dari tipe data yg sudah ada agar lebih mudah dibaca dan dikelola
	type Married bool
	
	var noKtpFadli NoKTP = "1928374618301837"
	var marriedStatus Married = false

	fmt.Println(noKtpFadli)
	fmt.Println(marriedStatus)

}