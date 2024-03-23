/*
Package string adalah package yg berisikan function-function untuk memanipulasi tipe data string
*/

package main

import (
	"fmt"
    "strings"
)

func main() {
	fmt.Println(strings.Contains("Fadli Darusalam", "Fadli"))
	fmt.Println(strings.Contains("Fadli Darusalam", "Joko"))
	fmt.Println()
	fmt.Println(strings.Split("Fadli Darusalam", " "))
	fmt.Println()
	fmt.Println(strings.ToLower("Fadli Darusalam"))
	fmt.Println(strings.ToUpper("Fadli Darusalam"))
	fmt.Println(strings.ToTitle("Fadli Darusalam"))
	fmt.Println()
	fmt.Println(strings.Trim("   Fadli Darusalam   ", " "))
	fmt.Println()
	fmt.Println(strings.ReplaceAll("Fadli Darusalam Fadli", "Fadli", "Bolt"))
}