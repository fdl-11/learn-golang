/*
[=] Package [=]
Package adalah tempat yg bisa digunakan untuk mengorganisir kode program yg telah kita buat
Dengan menggunakan package, kita bisa merapikan kode program yg kita buat
Package sendiri sebenarnya hanya direktori folder di sistem operasi kita

[=] Import [=]
Secara standar, file Golang hanya bisa mengakses file golang lainnya yg berada dalam package yg sama
Jika kita ingin mengakses file Golang yg berada di luar package, maka kita bisa menggunakan import
*/

package helper

import "fmt"

// Tidak bisa diakses dari luar package
var version = 1

// Bisa diakses dari luar package
var Application = "Belajar Golang"

// Bisa diakses dari luar package
func SayHello(name string) {
	fmt.Println("Hello " + name)
}

// Tidak bisa diakses dari luar packag
func sayGoodBye(name string) {
	fmt.Println("Good Bye " + name)
}