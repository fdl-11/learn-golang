/*
Saat membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
Maka, jika kita mengubah data dalam function, data yg asli tidak akan pernah berubah

Namun, kadang kita ingin membuat function yg bisa mengubah data asli dari parameter tersebut.
Untuk melakukan ini, kita bisa menggunakan pointer di function
Untuk menjadikan sebuah parameter menjadi pointer, kita bisa menggunakan operator "*" di parameternya

"Kalau membuat struct dengan data/field yg banyak atau besar, lebih baik menggunakan pointer agar tidak membebani memori"
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Sragen", "Jawa Tengah", "Indonesia"}
	address2 := &address1
	var address3 *Address = &address1
	
	address2.City = "Lampung"
	
	*address2 = Address{"Kuningan", "Jawa Barat", "Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	
	fmt.Println()

	var address4 *Address = new(Address)
	address4.City = "Solo"
	fmt.Println(address4)

	fmt.Println()

	var alamat = Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "",
	}
	// var alamatPointer *Address = &alamat
	// ChangeCountryToIndonesia(alamatPointer)
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}