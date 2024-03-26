/*
Package encoding berfungsi untuk melakukan encode dan decode
Algoritma yg paling populer adalah base64, csv, dan json
*/

package main

import (
	"fmt"
	"encoding/base64"
)

func main() {
	value := "Fadli Darusalam"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	fmt.Println()

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}