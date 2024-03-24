/*
Package container/ring adalah implementasi struktur data circular list
Circular list adalah struktur data ring, dimana di akhir element akan kembali ke element awal (HEAD)
*/

package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// var data *ring.Ring = ring.New(5)
	data := ring.New(5)
	
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// fmt.Println(*data)

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}