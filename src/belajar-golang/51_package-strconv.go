/*
Package untuk mengkonversi string, misal dari int ke string, dkk atau sebaliknya
*/

package main

import(
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")				// string ke boolean
	// boolean, err := strconv.ParseBool("salah")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)		// string ke int
	if err == nil {
        fmt.Println(number)
    } else {
        fmt.Println(err.Error())
    }

	fmt.Println()

	// Tipe data lain ke string
	// value := strconv.FormatInt(1000000, 10)	// 10 = base decimal
	value := strconv.FormatInt(1000000, 2)	// 2 = base binary
	// value := strconv.FormatInt(1000000, 8)	// 8 = base oktal
	// value := strconv.FormatInt(1000000, 16)	// 16 = base hexadecimal
	fmt.Println(value)

	fmt.Println()

	// string ke int tanpa base tertentu
	valueInt, _ := strconv.Atoi("20000000")
	fmt.Println(valueInt)

	// int ke string tanpa base tertentu
	valueStr:= strconv.Itoa(20000)
	fmt.Println(valueStr)
}