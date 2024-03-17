package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // Terjadi integer overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// =================================================================

	var name = "Fadli"
	var e byte = name[0] // byte = uint8, ketika ambil satu karakter, maka bentuknya byte
	var eString string = string(e)
	var x = string(name[2])	// sama saja

	fmt.Println(e)
	fmt.Println(eString)
	fmt.Println(x)
}
