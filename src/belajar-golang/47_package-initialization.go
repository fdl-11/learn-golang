/*
[=] Package Initialization [=]
Saat membuat package, kita bisa membuat sebuah function yg akan diakses/dijalankan ketika package kita diakses
Sangat cocok ketika misal package berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init

[=] Blank Identifier [=]
Kadang kita ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function di package
Secara default, Golang akan complain ketika ada package yg diimport tapi tidak digunakan
Untuk menangani hal itu, kita bisa menggunakan blank identifier "_" sebelum nama package ketika ingin melakukan import
*/

package main

import (
	// "belajar-golang/database"
	_ "belajar-golang/database"	// blank identifier
	// "fmt"
)

func main() {
	// result := database.GetDatabase()
	// fmt.Println(result)
}