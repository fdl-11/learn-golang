/*
Package path digunakan untuk memanipulasi data path seperti path di URL atau path di file system
Secara default, package path menggunakan slash sebagai karakter pathnya, oleh karena itu cocok untuk data URL

Namun jika ingin menggunakannya untuk memanipulasi path di file system, karena windows menggunakan backslash, maka khusus untuk File System, perlu menggunakan package path/filepath
*/

package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}