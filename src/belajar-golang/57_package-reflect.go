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
	Name string `required:"true" max:"10"`
}

type COntohLagi struct {
	Name string
	Description string
}

// validasi dengan reflection
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// return reflect.ValueOf(data).Field(i).Interface() != ""
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
                return false
            }
		}
	}
	return true
}

func main() {
	sample := Sample{"Fadli"}

	sampleType := reflect.TypeOf(sample)
	// var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	fmt.Println()
	
	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := COntohLagi{"Fadli", "Darusalam"} // true karena tidak ada tag
	fmt.Println(IsValid(contoh))
}