/*
Di dalam bahasa pemrograman, biasanya ada fitur reflection, dimana kita bisa melihat struktur kode kita saat aplikasi sedang berjalan
Hal ini bisa dilakukan di golang dengan package reflect
Reflection sangat berguna ketika ingin membuat library yg general sehingga mudah digunakan
*/

package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Fadli"}

	sampleType := reflect.TypeOf(sample)
	// var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
}