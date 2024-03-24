/*
Package container/list adalah implementasi struktur data double linked list di Golang
*/

package main

import(
	"fmt"
    "container/list"
)

func main() {
	data := list.New()
	data.PushBack("Fadli")
	data.PushBack("Darusalam")
	data.PushBack("Sragen")
	data.PushFront("Mr")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Front().Next().Next().Value)
	fmt.Println(data.Back().Value)
	
	fmt.Println()
	
	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	fmt.Println()

	for e := data.Front(); e!= nil; e = e.Next() {
        fmt.Println(e.Value)
    }

	fmt.Println()

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}