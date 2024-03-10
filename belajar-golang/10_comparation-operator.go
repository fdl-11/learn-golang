package main

import "fmt"

func main() {
    var name1 = "Fadli"
	var name2 = "Fadli"
	var name3 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	var result2 bool = name2 == name3
	fmt.Println(result2)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}