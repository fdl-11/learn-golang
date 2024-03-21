/*
[=] Pass by Value [=]
Secara default, semua variabel di golang itu di passing by value, bukan by reference
Artinya, jika kita mengirim sebuah variabel ke dalam function, method, atau variabel lain ,sebenarnya yg kita kirim adalah duplikasi valuenya

[=] Pointer [=]
Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yg sudah ada
Dengan pointer, kita bisa membuat pass by reference
Untuk membuat variabel dengan nilai pointer ke variabel lain, kita bisa menggunakan operator "&" diikuti nama variabel

Saat kita mengubah variabel pointer, maka yg berubah hanya variabel tersebut
Semua variabel yg mengacu ke data yg sama tidak akan berubah
Jika ingin merubah seluruh variabel yg mengacu ke data tersebut, kita bisa menggunakan operator "*"

[=] Function new [=]
Golang memiliki function new() untuk membuat pointer, namun function new() hanya mengembalikan pointer ke data kosong / tidak ada data awal
*/

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Sragen", "Jawa Tengah", "Indonesia"}
	// address2 := address1	// pass by value
	address2 := &address1	// pass by reference / pointer
	// var address2 *Address = &address1	// sama saja
	var address3 *Address = &address1
	
	address2.City = "Lampung"
	
	// address2 = &Address{"Kuningan", "Jawa Barat", "Indonesia"}	// address1 tidak berubah
	*address2 = Address{"Kuningan", "Jawa Barat", "Indonesia"}	// address1 / semua yg mengacu ke address 1 akan berubah
	
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	
	fmt.Println()
	
	// var address4 *Address = &Address{"Sragen", "Jawa Tengah", "Indonesia"}	// membuat pointer baru
	var address4 *Address = new(Address)	// Membuat pointer menggunakan new(), namun datanya kosong
	address4.City = "Solo"
	fmt.Println(address4)
}