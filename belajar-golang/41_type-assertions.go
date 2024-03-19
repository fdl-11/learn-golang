/*
Type Assertions merupakan kemampuan mengubah tipe data menjadi tipe data yg diinginkan
Fitur ini sering digunakan ketika kita bertemu dengan data interface kosong

Saat salah menggunakan type assertions, maka berakibat panic di aplikasi kita
Jika panic dan tidak ada recover, maka otomatis program kita akan mati
Agar lebih aman, sebaiknya menggunakan bantuan switch expression untuk melakukan type assertions
*/

package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Value", value, "is unknown")
	}
}