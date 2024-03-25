/*
Package sort adalah package yg berisikan utilitas untuk proses pengurutan
Agar data bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.interface
*/

package main

import (
	"sort"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User {
		{"Fadli", 21},
		{"Cahya", 23},
		{"Haikal", 22}, 
		{"Nanda", 24},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}