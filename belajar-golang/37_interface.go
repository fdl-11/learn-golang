/**
Interface adalah tipe data Abstract, tidak memiliki implementasi langsung
Interface berisi definisi-definisi method
Biasanya interface digunakan sebagai definisi kontrak, untuk function-function yang bersifat general

Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
Sehingga kita tidak perlu mengimplementasikan interface secara manual
*/

package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	fadli := Person{Name: "Fadli"}
	SayHello(fadli)

	fmt.Println()

	cat := Animal{Name: "Kitty"}
	SayHello(cat)
}