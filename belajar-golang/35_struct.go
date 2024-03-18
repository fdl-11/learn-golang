/**
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
Struct dapat disebut juga kumpulan dari field
Struct tidak bisa langsung digunakan, namun bisa membuat data/objek dari struct yang telah dibuat
*/

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main() {
	var fadli Customer
	fadli.Name = "Fadli"
	fadli.Address = "Sragen"
	fadli.Age = 21

	fmt.Println(fadli)
	fmt.Println(fadli.Name)
	fmt.Println(fadli.Address)
	fmt.Println(fadli.Age)

	fmt.Println()
	
	darusalam := Customer{
		Name: "Darusalam",
        Address: "Sragen",
        Age: 21,
	}
	fmt.Println(darusalam)
	
	fmt.Println()

	joko := Customer{"Joko", "Solo", 31}	// tidak disarankan
	fmt.Println(joko)

}