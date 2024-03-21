/*
Walaupun method akan menempel di struct, tapi sebenarnya data struct yg diakses di dalam method adalah pass by value
Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memori karena harus diduplikasi ketika memanggil method
*/

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	fadli := Man{"Fadli"}
	fadli.Married()

	fmt.Println(fadli)
}