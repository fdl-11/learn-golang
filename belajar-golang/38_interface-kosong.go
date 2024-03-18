/**
Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya
- fmt.Println(a ..interface[])
- panic(v interface[])
- recover() interface[]

misal kita ingin membuat semua tipe data ingin diterima di parameter tersebut
*/

package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(1)
	fmt.Println(data)
}