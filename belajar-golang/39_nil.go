/*
Data kosong di golang disebut nil
Nil hanya bisa digunakan di beberapa tipe data seperti interface, function, map, slice, pointer, dan channel
*/

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Fadli")
	fmt.Println(person)

	fmt.Println()

	var person2 map[string]string = NewMap("")
	if person2 == nil {
		fmt.Println("Data Kosong!")
	} else {
		fmt.Println(person2)
	}

}