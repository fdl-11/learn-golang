/*
Di golang terdapat fitur Generic, dimana kita bisa membuat parameter dengan tipe yang bisa berubah-ubah tanpa menggunakan interface kosong/any
Salah satu package yang menggunakan fitur ini adalah package slices, untuk memanipulasi data di slice
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Fadli", "Joko", "Budi", "Adi"}
	values := []int{100, 90, 92, 88}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Fadli"))
	fmt.Println(slices.Index(names, "Fadli"))
	fmt.Println(slices.Index(names, "Eko"))
}